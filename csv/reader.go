package csv

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

type Csv struct {
	filename string
	file     *os.File
	r        *csv.Reader
	fields   []string
	values   []string
	eof      bool
	err      error
}

func NewCsvReader(filename string, includeHeader bool) (*Csv, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)
	r.TrimLeadingSpace = true
	r.LazyQuotes = true

	var fields []string
	if includeHeader {
		records, err := r.Read()
		if err != nil {
			return nil, err
		}
		fields = records
		for i := 0; i < len(fields); i++ {
			fields[i] = strings.ReplaceAll(fields[i], `"`, "")
			fields[i] = removeControlRune(fields[i])
		}
	}

	r.ReuseRecord = true
	return &Csv{
		filename: filename,
		file:     file,
		r:        r,
		fields:   fields,
	}, nil
}

func (s *Csv) SetHeader(header []string) {
	s.fields = header
}

func (s *Csv) Name() string {
	return s.filename
}

func (s *Csv) Next() bool {
	if s.eof || s.err != nil {
		return false
	}

	values, err := s.r.Read()
	if err != nil {
		if err == io.EOF {
			s.eof = true
		} else {
			s.err = err
		}
		return false
	}

	s.values = values
	return true
}

func (s *Csv) Get() (*Record, error) {
	r := &Record{KV: map[string]string{}}
	for i, field := range s.fields {
		r.KV[field] = s.values[i]
	}
	return r, nil
}

func (s *Csv) GetRecord() ([]string, error) {
	return s.values, nil
}

func (s *Csv) Scan(obj interface{}) error {
	r, err := s.Get()
	if err != nil {
		return err
	}

	return scan(obj, r)
}

func (s *Csv) ScanRecord(r *Record) error {
	for i, field := range s.fields {
		r.KV[field] = s.values[i]
	}
	return nil
}

func (s *Csv) Error() error {
	return s.err
}

func (s *Csv) Close() error {
	return s.file.Close()
}

func removeControlRune(s string) string {
	ret := ""
	for _, r := range s {
		if r == 65279 {
			continue
		}
		ret += string(r)
	}
	return ret
}
