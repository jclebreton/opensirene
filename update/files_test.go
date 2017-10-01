package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getScratchZipList(t *testing.T) {
	url := ""
	dest := ""
	zipList, err := getScratchZipList("Jan", url, dest)
	assert.NoError(t, err)
	assert.True(t, len(zipList) > 0)
	assert.Contains(t, zipList[0].name, "sirene_")
}

func Test_getIncrementalFile(t *testing.T) {
	file := getIncrementalFile(time.Now(), "/tmp", "http://example.com")
	assert.Regexp(t, "sirene_[0-9]+_E_Q", file.name)
	assert.Regexp(t, "/tmp/sirene_[0-9]{7}_E_Q.zip", file.path)
}

func Test_getCompleteFile(t *testing.T) {
	file := getCompleteFile(time.Now(), "/tmp", "http://example.com")
	assert.Regexp(t, "sirene_[0-9]{6}_L_M", file.name)
	assert.Regexp(t, "/tmp/sirene_[0-9]{6}_L_M.zip", file.path)
}

func Test_getFile(t *testing.T) {
	file := getFile("foo", "complete", "/tmp", "http://example.com")
	assert.Equal(t, "foo", file.name)
	assert.Equal(t, "complete", file.updateType)
	assert.Equal(t, "foo.zip", file.filename)
	assert.Equal(t, "http://example.com/foo.zip", file.url)
	assert.Equal(t, "/tmp/foo.zip", file.path)
}
