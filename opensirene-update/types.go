package main

import "io"

type zipFile struct {
	name       string
	filename   string
	path       string
	url        string
	updateType string
	csvFiles   []csvFile
}

type csvFile struct {
	filename string
	path     string
}

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
