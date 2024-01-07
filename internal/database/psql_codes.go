package database

import (
	"errors"
	"regexp"
	"strings"
)

// SQLSTATE described in: https://www.postgresql.org/docs/current/errcodes-appendix.html
const (
	cannot_connect_now_code = "57P03"
)

var (
	errCannotConnectNow = errors.New("unable to connect to database now")
	rxSQLStateError     = regexp.MustCompile("SQLSTATE [A-Z0-9]{5}")
)

func getSQLErrorCode(err error) string {
	s := rxSQLStateError.FindString(err.Error())
	if s == "" {
		return ""
	}

	return strings.Split(s, " ")[1]
}
