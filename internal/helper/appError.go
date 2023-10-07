package helper

import (
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

//func (e *AppError) RootError() error {
//	if err, ok := e.RootErr.(*AppError); ok {
//		return err.RootError()
//	}
//
//	return e.RootErr
//}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string, statusCode ...int) *AppError {
	if statusCode != nil {
		return NewFullErrorResponse(statusCode[0], root, msg, root.Error(), key)
	}

	return NewErrorResponse(root, msg, root.Error(), key)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "internal error", err.Error(), "ErrInternal")
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannnotUpdate%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrEntityDeleted(entity string) *AppError {
	return NewCustomError(
		fmt.Errorf("%s invalid", strings.ToLower(entity)),
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

func ErrCannotAddItemToEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot add item to %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotAddItemTo%s", entity),
	)
}

func ErrCannotRemoveItemFromEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot remove item from %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotRemoveItemFrom%s", entity),
	)
}

//func ErrDbb(_err error) error {
//	if _err == nil {
//		return nil
//	}
//
//	switch _err {
//	case gorm.ErrRecordNotFound:
//		return errors.New("record not found")
//	default:
//		var mysqlErr *mysql.MySQLError
//		if errors.As(_err, &mysqlErr) && mysqlErr.Number == 1062 {
//			return errors.New(extractColumnName(mysqlErr.Message) + " already exists.")
//		}
//	}
//	return nil //----
//	//return NewErrorResponse(err, "Something went wrong with DB", err.Error(), "DB_ERROR")
//}

func ErrDb(err error) *AppError {
	return NewErrorResponse(err, "Something went wrong with DB", err.Error(), "DB_ERROR")
}

//func extractColumnName(text string) string {
//	re := regexp.MustCompile(`'([^']+)'\s+for\s+key\s+'([^']+)'`)
//	if re.MatchString(text) {
//		return re.FindStringSubmatch(text)[2]
//	}
//	return "Unknown"
//}
