package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/implicithash/bookstore_users-api/utils/errors"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NotFoundError("no error matching given id")
		}
		return errors.InternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError("invalid data")
	}
	return errors.InternalServerError("error processing request")
}
