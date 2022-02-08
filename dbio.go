package db

type Scanner interface {
	Next() bool
	MapScan(dest map[string]interface{}) error
	Close() error
}
