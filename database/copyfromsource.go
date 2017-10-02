package database

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"

	iconv "github.com/djimenez/iconv-go"
)

type source struct {
	file     *os.File
	path     string
	scanner  *bufio.Scanner
	values   []interface{}
	err      error
	progress chan<- map[string]float64
	cpt      float64
	total    float64
}

func InitCopyFromFile(path string, progress chan<- map[string]float64) (*source, error) {
	source := &source{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	source.file = file
	source.path = path
	source.scanner = bufio.NewScanner(file)
	source.progress = progress

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	source.total = float64(stat.Size())
	source.Next()

	return source, nil
}

// Next returns true if there is another row and makes the next row data
// available to Values(). When there are no more rows available or an error
// has occurred it returns false.
func (s *source) Next() bool {

	//Read line
	ok := s.scanner.Scan()

	//EOF
	if !ok {
		s.progress <- map[string]float64{s.path: 100}
		return false
	}

	line := s.scanner.Text()
	err := s.scanner.Err()
	if err != nil {
		defer s.file.Close()
		s.progress <- map[string]float64{s.path: 100}
		s.err = err
		return false
	}

	//Convert to UTF8
	lineUTF8, _ := iconv.ConvertString(line, "windows-1252", "utf-8")

	//Parse line
	r := csv.NewReader(strings.NewReader(lineUTF8))
	r.Comma = ';'
	records, err := r.Read()
	if err != nil {
		s.progress <- map[string]float64{s.path: 100}
		s.err = err
		return false
	}

	//Build result
	var values []interface{}
	for _, v := range records {
		values = append(values, v)
	}
	s.values = values

	//Progress
	s.cpt += float64(len(line))
	s.progress <- map[string]float64{s.path: (s.cpt / s.total) * 100}

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
