package hive

import (
	"context"
	"database/sql"

	"github.com/beltran/gohive"
	db "github.com/boram-gong/db_tool"
)

type HiveDB struct {
	con *gohive.Connection
}

func NewHiveClient(hiveCfg *db.CfgDB) db.DB {
	conCfg := gohive.NewConnectConfiguration()
	conCfg.Database = hiveCfg.Database
	conCfg.HiveConfiguration = make(map[string]string)
	conCfg.FetchSize = 1000
	conCfg.Service = "hive"
	conCfg.Username = hiveCfg.User
	conCfg.Password = hiveCfg.Password
	conn, err := gohive.Connect(hiveCfg.Host, hiveCfg.Port, "NONE", conCfg)
	if err != nil {
		panic(err)
	}
	return &HiveDB{con: conn}
}

func (db *HiveDB) QueryX(query string, args ...interface{}) (db.Scanner, error) {
	if len(args) > 0 {
		query = replace(query, args)
	}

	ctx := context.Background()
	cursor := db.con.Cursor()
	cursor.Exec(ctx, query)
	if cursor.Err != nil {
		return nil, cursor.Err
	}
	desc := cursor.Description()
	fields := make([]string, 0, len(desc))
	for _, field := range desc {
		fields = append(fields, field[0])
	}
	return &HiveRows{ctx: ctx, fields: fields, cursor: cursor}, nil
}

func (db *HiveDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if len(args) > 0 {
		query = replace(query, args)
	}

	cursor := db.con.Cursor()
	cursor.Exec(context.Background(), query)
	return nil, cursor.Err
}

type HiveRows struct {
	ctx    context.Context
	fields []string
	cursor *gohive.Cursor
}

func (r *HiveRows) Next() bool {
	return r.cursor.HasMore(r.ctx)
}

func (r *HiveRows) MapScan(dest map[string]interface{}) error {
	results := make([]interface{}, len(r.fields))
	r.cursor.FetchOne(r.ctx, results...)
	if r.cursor.Err != nil {
		return r.cursor.Err
	}

	for i, field := range r.fields {
		dest[field] = results[i]
	}
	return nil
}

func (r *HiveRows) Close() error {
	r.cursor.Close()
	return r.cursor.Err
}

func replace(query string, args ...interface{}) string {
	panic("not implement")
}
