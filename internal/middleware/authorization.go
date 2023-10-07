package middleware

import (
	"ecommercestore/internal/conf"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/helper/jwt"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

var ErrWrongAuthHeader = helper.NewCustomError(
	errors.New("wrong authentication header"),
	"wrong authentication header",
	"ErrWrongAuthHeader",
	http.StatusUnauthorized,
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 || parts[0] != "Bearer" || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader
	}
	return parts[1], nil
}

func Authorization() func(c *gin.Context) {
	tokenProvider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	return func(ctx *gin.Context) {
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))
		if err != nil {
			log.Error().Err(err).Msg("Error extractTokenFromHeaderString")
			panic(err)
		}

		id, err := tokenProvider.ValidateToken(token)
		if err != nil {
			panic(err)
		}
		user, err := repo.GetUser(map[string]interface{}{"id": id})
		if err != nil {
			panic(err)
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

func CurrentUser(ctx *gin.Context) (*models.User, error) {
	var err error
	_user, exists := ctx.Get("user")
	if !exists {
		err = errors.New("current context user not set")
		log.Error().Err(err).Msg("")
		return nil, err
	}
	user, ok := _user.(*models.User)
	if !ok {
		err = errors.New("context user is not valid type")
		log.Error().Err(err).Msg("")
		return nil, err
	}
	return user, nil
}
