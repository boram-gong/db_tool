package mysql

import (
	"fmt"
	"testing"

	db "github.com/boram-gong/db_tool"
)

func TestMysql(t *testing.T) {
	cfg := &db.CfgDB{
		Host:        "114.67.66.0",
		Port:        3456,
		User:        "root",
		Password:    "1qazZSE$",
		Database:    "chengdu",
		MaxIdleConn: 2,
		MaxOpenConn: 100,
	}
	cli := NewMysqlClient(cfg)
	querySql := "select Pkid,FmContent from export where DCntAName='违章建筑'or DCntAName='占道经营' limit 10;"
	rows, err := cli.QueryX(querySql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		m := map[string]interface{}{}
		if err := rows.MapScan(m); err != nil {
			continue
		}
		fmt.Println(string(m["Pkid"].([]uint8)))
	}
}
