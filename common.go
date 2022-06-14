package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/boram-gong/db_tool/mysql"
	"github.com/boram-gong/db_tool/pg"
	"strconv"
)

type CfgDB struct {
	Host        string `yaml:"Host"`
	Port        int    `yaml:"Port"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Database    string `yaml:"Database"`
	MaxIdleConn int    `yaml:"MaxIdleConn"`
	MaxOpenConn int    `yaml:"MaxOpenConn"`
}

func NewDbClient(dbType string, cfg *CfgDB) (DB, error) {
	switch dbType {
	case "postgres":
		return pg.NewPgClient(cfg)
	case "mysql":
		return mysql.NewMysqlClient(cfg)
	}
	return nil, errors.New("db type err")
}

type DB interface {
	QueryX(query string, args ...interface{}) (Scanner, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Scanner interface {
	Next() bool
	MapScan(dest map[string]interface{}) error
	Close() error
}

func Interface2Map(data interface{}) map[string]interface{} {
	switch data.(type) {
	case map[string]interface{}:
		return data.(map[string]interface{})
	default:
		return nil
	}
}

func Interface2Slice(data interface{}) []interface{} {
	switch data.(type) {
	case []interface{}:
		return data.([]interface{})
	default:
		return nil
	}
}

func Interface2String(data interface{}) string {
	switch data.(type) {
	case nil:
		return ""
	case string:
		return data.(string)
	case []byte:
		return string(data.([]byte))
	default:
		return fmt.Sprintf("%v", data)
	}
}

func Interface2Int(data interface{}) int {
	switch data.(type) {
	case int64:
		return int(data.(int64))
	case int:
		return data.(int)
	case string:
		i, _ := strconv.Atoi(data.(string))
		return i
	case float64:
		return int(data.(float64))
	default:
		return 0
	}
}

func Interface2Float(data interface{}) float64 {
	switch data.(type) {
	case float64:
		return data.(float64)
	case int64:
		return float64(data.(int64))
	case int:
		return float64(data.(int))
	case string:
		i, _ := strconv.ParseFloat(data.(string), 64)
		return i
	default:
		return 0.0
	}
}
