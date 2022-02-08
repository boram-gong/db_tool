package csv

import (
	"encoding/csv"
	"os"
)

func NewCsvWriter(filename string) (*csv.Writer, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return csv.NewWriter(file), nil
}
