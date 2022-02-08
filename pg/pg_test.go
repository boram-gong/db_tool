package pg

import (
	db "db_tool"
	"fmt"
	"testing"
)

func TestMysql(t *testing.T) {
	cfg := &db.CfgDB{
		Host:        "172.2.0.21",
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
