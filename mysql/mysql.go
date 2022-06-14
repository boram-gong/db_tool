package mysql

import (
	"database/sql"
	"fmt"
	"github.com/boram-gong/db_tool/common"
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlClient(cfg *common.CfgDB) (*MClient, error) {
	cli, err := NewMysql(cfg)
	if err != nil {
		return nil, err
	}
	return &MClient{cli}, nil
}

func NewMysql(cfg *common.CfgDB) (*sqlx.DB, error) {
	dbURI := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?loc=Local&parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	con, err := sqlx.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	// 配置连接池
	con.SetConnMaxLifetime(time.Second * 5)
	con.SetMaxIdleConns(cfg.MaxIdleConn)
	con.SetMaxOpenConns(cfg.MaxOpenConn)
	err = con.Ping()
	if err != nil {
		return nil, err
	}
	return con, nil
}

type MClient struct {
	mClient *sqlx.DB
}

func (cli *MClient) QueryX(query string, args ...interface{}) (common.Scanner, error) {
	return cli.mClient.Queryx(query, args...)
}

func (cli *MClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return cli.mClient.Exec(query, args...)
}
