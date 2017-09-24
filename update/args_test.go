package main

import (
	"testing"

	"os"

	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getWorkingDirectory_manual(t *testing.T) {
	args := make(map[string]interface{})
	args["--wd"] = "/tmp/foo"
	wd, err := getWorkingDirectory(args)
	assert.NoError(t, err)
	assert.Equal(t, "/tmp/foo", wd)
	stat, err := os.Stat("/tmp/foo")
	assert.NoError(t, err)
	assert.True(t, stat.IsDir())
}

func Test_getWorkingDirectory_auto(t *testing.T) {
	args := make(map[string]interface{})
	wd, err := getWorkingDirectory(args)
	assert.NoError(t, err)
	stat, err := os.Stat(wd)
	assert.NoError(t, err)
	assert.True(t, stat.IsDir())
}

func Test_getWorkingDirectory_error(t *testing.T) {
	args := make(map[string]interface{})
	args["--wd"] = "/root"
	_, err := getWorkingDirectory(args)
	assert.Error(t, err)
}

func Test_getNbWorkers_manual(t *testing.T) {
	args := make(map[string]interface{})
	args["--maxworkers"] = "2"
	max, err := getNbWorkers(args)
	assert.NoError(t, err)
	assert.Equal(t, 2, max)
}

func Test_getNbWorkers_auto(t *testing.T) {
	args := make(map[string]interface{})
	max, err := getNbWorkers(args)
	assert.NoError(t, err)
	assert.Equal(t, nbWorkersMax, max)
}

func Test_getMonth_manual(t *testing.T) {
	args := make(map[string]interface{})
	args["--month"] = "Aug"
	month := getMonth(args)
	assert.Equal(t, "Aug", month)
}

func Test_getMonth_auto(t *testing.T) {
	args := make(map[string]interface{})
	month := getMonth(args)
	assert.Equal(t, time.Now().Format("Jan"), month)
}
