package db

import (
	"database/sql"
	"errors"
	"github.com/boram-gong/db_tool/common"
	"github.com/boram-gong/db_tool/hive"
	"github.com/boram-gong/db_tool/mysql"
	"github.com/boram-gong/db_tool/pg"
)

func NewDbClient(dbType string, cfg *common.CfgDB) (DB, error) {
	switch dbType {
	case "postgres":
		return pg.NewPgClient(cfg)
	case "mysql":
		return mysql.NewMysqlClient(cfg)
	case "hive":
		return hive.NewHiveClient(cfg)

	}
	return nil, errors.New("db type err")
}

type DB interface {
	QueryX(query string, args ...interface{}) (common.Scanner, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}
