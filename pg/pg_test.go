package pg

import (
	"fmt"
	"testing"

	db "github.com/boram-gong/db_tool"
)

func TestMysql(t *testing.T) {
	cfg := &db.CfgDB{
		Host:        "172.2.0.0",
		Port:        5436,
		User:        "test_gyx",
		Password:    "wayzpg_gyx",
		Database:    "pro_dimingdizhixiangmu",
		MaxIdleConn: 2,
		MaxOpenConn: 100,
	}
	cli := NewPgClient(cfg)
	querySql := "select * from address_standardize limit 10;"
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
		fmt.Println(m)
	}
}
