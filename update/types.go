package main

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
