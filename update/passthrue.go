package main

import "io"

// PassThru code originally from
// http://stackoverflow.com/a/22422650/613575
type PassThru struct {
	io.Reader
	curr     int64
	total    float64
	filename string
	progress chan map[string]float64
}

//Override native Read
func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	pt.curr += int64(n)

	// last read will have EOF err
	if err == nil || (err == io.EOF && n > 0) {
		pt.progress <- map[string]float64{pt.filename: (float64(pt.curr) / pt.total) * 100}
	}

	return n, err
}
