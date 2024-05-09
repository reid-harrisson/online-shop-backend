package errhandle

import (
	"OnlineStoreBackend/pkgs/constants"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func SqlErrorHandler(err error) (int, string) {
	sqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		switch sqlErr.Number {
		case constants.ErrInvalidData:
			return http.StatusBadRequest, constants.InvalidData
		case constants.ErrInvalidValueOfLength:
			return http.StatusBadRequest, constants.InvalidData
		case constants.ErrDuplicatedKeyRequired:
			return http.StatusBadRequest, constants.InvalidData
		case constants.ErrForeignKeyViolated:
			return http.StatusBadRequest, constants.InvalidData
		case constants.ErrInvalidField:
			return http.StatusInternalServerError, constants.InternalServerErrorMessage
		case constants.ErrRecordNotFound:
			return http.StatusNotFound, constants.NotFound
		case constants.ErrMissingWhereClause:
			return http.StatusInternalServerError, constants.InternalServerErrorMessage
		}
	} else {
		switch err {
		case gorm.ErrEmptySlice:
			return http.StatusInternalServerError, constants.InternalServerErrorMessage
		case gorm.ErrRecordNotFound:
			return http.StatusNotFound, constants.NotFound
		case gorm.ErrForeignKeyViolated:
			return http.StatusBadRequest, constants.InvalidData
		}
	}
	if err != nil {
		return http.StatusInternalServerError, constants.InternalServerErrorMessage
	}
	return 0, ""
}
