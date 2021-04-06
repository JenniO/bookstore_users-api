package mysql_utils

import (
	"fmt"
	"github.com/JenniO/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError(fmt.Sprintf("error parsing database response: %s", err.Error()))
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(sqlErr.Message)
	}
	return errors.NewInternalServerError(fmt.Sprintf("error parsing database response: %s", err.Error()))
}
