package database

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"
)

type source struct {
	file    *os.File
	scanner *bufio.Scanner
	values  []interface{}
	err     error
	cpt     int
}

func InitCopyFromSource(path string) *source {
	source := &source{}

	file, err := os.Open(path)
	if err != nil {
		source.err = err
	}

	source.file = file
	source.scanner = bufio.NewScanner(file)
	source.Next()
	source.cpt = 0
	return source
}

// Next returns true if there is another row and makes the next row data
// available to Values(). When there are no more rows available or an error
// has occurred it returns false.
func (s *source) Next() bool {
	s.scanner.Scan()
	line := s.scanner.Text()

	err := s.scanner.Err()
	if err != nil {
		defer s.file.Close()
		s.err = err
		return false
	}

	r := csv.NewReader(strings.NewReader(line))
	r.Comma = ';'

	records, err := r.Read()
	if err != nil {
		s.err = err
		return false
	}

	var values []interface{}
	for _, v := range records {
		values = append(values, v)
	}
	s.values = values
	if s.cpt == 1000 {
		return false
	}

	s.cpt++
	return true
}

// Values returns the values for the current row.
func (s *source) Values() ([]interface{}, error) {
	return s.values, s.err
}

// Err returns any error that has been encountered by the CopyFromSource. If
// this is not nil *Conn.CopyFrom will abort the copy.
func (s *source) Err() error {
	return s.err
}
