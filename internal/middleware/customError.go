package middleware

import (
	"ecommercestore/internal/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

func CustomErrors(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		ctx.Header("Content-Type", "application/json")
		for _, err := range ctx.Errors {
			switch err.Type {
			case gin.ErrorTypeBind:
				errMap := make(map[string]string)
				if errs, ok := err.Err.(validator.ValidationErrors); ok {
					for _, fieldErr := range []validator.FieldError(errs) {
						errMap[fieldErr.Field()] = customValidationError(fieldErr)
					}
				}
				status := http.StatusBadRequest
				if ctx.Writer.Status() != http.StatusOK {
					status = ctx.Writer.Status()
				}
				ctx.AbortWithStatusJSON(status, gin.H{"error": errMap})
			default:
				// Log other errors
				log.Error().Err(err.Err).Msg("Other error")
			}
		}
		if !ctx.Writer.Written() {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError"})
		}
	}
}

func customValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required.", err.Field())
	case "min":
		return fmt.Sprintf("%s must be longer than or equal %s characters.", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s characters.", err.Field(), err.Param())
	case "email":
		return fmt.Sprintf("%s must be in correct format.", err.Field())
	default:
		return err.Error()
	}
}

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.Header("Content-Type", "application/json")

			if appErr, ok := err.(*helper.AppError); ok {
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				//panic(err)
				return
			}

			appErr := helper.ErrInternal(err.(error))
			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			//panic(err)
			return
		}
	}()

	c.Next()
}
