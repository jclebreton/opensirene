package main

import (
	"io/ioutil"
	"os"

	docopt "github.com/docopt/docopt-go"

	"strconv"

	"errors"

	"time"

	"golang.org/x/sys/unix"
)

type args struct {
	list map[string]interface{}
}

func InitArgs(usage string) args {
	arguments, _ := docopt.Parse(usage, nil, true, "", false)
	return args{list: arguments}
}

func (a *args) getWorkingDirectory() (string, error) {

	var wd string

	if a.list["--wd"] != nil {
		wd = a.list["--wd"].(string)
	} else {
		temp, err := ioutil.TempDir("/tmp/", "tmp")
		if err != nil {
			return "", err
		}
		wd = temp
	}

	err := os.MkdirAll(wd, os.ModePerm)
	if err != nil {
		return "", err
	}

	if unix.Access(wd, unix.W_OK) != nil {
		return "", errors.New("Working directory is not writable")
	}

	return wd, nil
}

func (a *args) getNbWorkers() (int, error) {
	if a.list["--maxworkers"] != nil {
		n, err := strconv.Atoi(a.list["--maxworkers"].(string))
		if err != nil || n <= 0 || n > nbWorkersMax {
			return 0, errors.New("Number of workers must be >= 1 and <=31")
		}
		return n, nil
	}

	return nbWorkersMax, nil
}

func (a *args) getMonth() string {
	var month string

	if a.list["--month"] != nil {
		month = a.list["--month"].(string)
	} else {
		month = time.Now().Format("Jan")
	}

	return month
}

func (a *args) isDebug() bool {
	return a.list["--debug"].(bool)
}

func (a *args) isCompleteUpdate() bool {
	return a.list["complete"].(bool)
}

func (a *args) isDailyUpdate() bool {
	return a.list["daily"].(bool)
}
