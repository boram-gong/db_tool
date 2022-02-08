package csv

import (
	"reflect"
	"strings"
)

const (
	tagCsv = "csv"
)

type Source interface {
	Name() string
	Next() bool
	Get() (*Record, error)
	GetRecord() ([]string, error)
	Scan(obj interface{}) error
	ScanRecord(record *Record) error
	Error() error
	Close() error
}

type Record struct {
	KV map[string]string
}

func scan(obj interface{}, r *Record) error {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag

		key, ok := tag.Lookup(tagCsv)
		if !ok {
			key = field.Name
		}

		value, ok := r.KV[key]
		if !ok {
			key = toLower(key)
		}

		value, ok = r.KV[key]
		if !ok {
			continue
		}

		v.Field(i).SetString(value)
	}
	return nil
}

func toLower(s string) string {
	rns := []rune(s)
	return strings.ToLower(string(rns[0])) + string(rns[1:])
}
