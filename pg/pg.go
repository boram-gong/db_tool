package pg

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/boram-gong/db_tool/common"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPgClient(pgCfg *common.CfgDB) (*PClient, error) {
	cli, err := NewPostgreDB(pgCfg)
	if err != nil {
		return nil, err
	}
	return &PClient{cli}, nil
}

func NewPostgreDB(cfg *common.CfgDB) (*sqlx.DB, error) {
	dbURI := fmt.Sprintf("user=%v password=%v sslmode=disable dbname=%v host=%v port=%v",
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.Host,
		cfg.Port,
	)
	con, err := sqlx.Open("postgres", dbURI)
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

type PClient struct {
	pClient *sqlx.DB
}

func (cli *PClient) QueryX(query string, args ...interface{}) (common.Scanner, error) {
	return cli.pClient.Queryx(query, args...)
}

func (cli *PClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return cli.pClient.Exec(query, args...)
}
