package main

import (
	"testing"

	"os"

	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getWorkingDirectory_manual(t *testing.T) {
	list := make(map[string]interface{})
	list["--wd"] = "/tmp/foo"
	args := args{list: list}
	wd, err := args.getWorkingDirectory()
	assert.NoError(t, err)
	assert.Equal(t, "/tmp/foo", wd)
	stat, err := os.Stat("/tmp/foo")
	assert.NoError(t, err)
	assert.True(t, stat.IsDir())
}

func Test_getWorkingDirectory_auto(t *testing.T) {
	args := args{}
	wd, err := args.getWorkingDirectory()
	assert.NoError(t, err)
	stat, err := os.Stat(wd)
	assert.NoError(t, err)
	assert.True(t, stat.IsDir())
}

func Test_getWorkingDirectory_error(t *testing.T) {
	list := make(map[string]interface{})
	list["--wd"] = "/root"
	args := args{list: list}
	_, err := args.getWorkingDirectory()
	assert.Error(t, err)
}

func Test_getNbWorkers_manual(t *testing.T) {
	list := make(map[string]interface{})
	list["--maxworkers"] = "2"
	args := args{list: list}
	max, err := args.getNbWorkers()
	assert.NoError(t, err)
	assert.Equal(t, 2, max)
}

func Test_getNbWorkers_auto(t *testing.T) {
	list := make(map[string]interface{})
	args := args{list: list}
	max, err := args.getNbWorkers()
	assert.NoError(t, err)
	assert.Equal(t, nbWorkersMax, max)
}

func Test_getMonth_manual(t *testing.T) {
	list := make(map[string]interface{})
	list["--month"] = "Aug"
	args := args{list: list}
	month := args.getMonth()
	assert.Equal(t, "Aug", month)
}

func Test_getMonth_auto(t *testing.T) {
	args := args{}
	month := args.getMonth()
	assert.Equal(t, time.Now().Format("Jan"), month)
}
