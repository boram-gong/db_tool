package db

import (
	"database/sql"
)

type DB interface {
	QueryX(query string, args ...interface{}) (Scanner, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}
