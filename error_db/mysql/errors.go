package mysql

import (
	errno "github.com/bombsimon/mysql-error-numbers"
	"github.com/go-sql-driver/mysql"
)

// MySQLError mysql database error
type MySQLError errno.ErrorNumber

func (e MySQLError) Error() string {
	return errno.ErrorNumber(e).String()
}

// ParseMySQLError convert mysql error to MySQLError, or else returns original error.
func ParseMySQLError(err error) error {
	if _, ok := err.(*mysql.MySQLError); ok {
		return MySQLError(errno.FromError(err))
	}
	return err
}

// IsMySQLError check is a mysql error number
func IsMySQLError(err error, dbErr errno.ErrorNumber) bool {
	if _, ok := err.(*mysql.MySQLError); ok {
		return errno.FromError(err) == dbErr
	} else if ex, ok := err.(MySQLError); ok {
		return errno.ErrorNumber(ex) == dbErr
	}
	return false
}
