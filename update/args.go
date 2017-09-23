package main

import (
	"io/ioutil"
	"os"

	"strconv"

	"errors"

	"time"

	"golang.org/x/sys/unix"
)

func getWorkingDirectory(arguments map[string]interface{}) (string, error) {

	var wd string

	if arguments["--wd"] != nil {
		wd = arguments["--wd"].(string)
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

func getNbWorkers(arguments map[string]interface{}) (int, error) {
	if arguments["--maxworkers"] != nil {
		n, err := strconv.Atoi(arguments["--maxworkers"].(string))
		if err != nil || n <= 0 || n > nbWorkersMax {
			return 0, errors.New("Number of workers must be >= 1 and <=100")
		}
		return n, nil
	}

	return nbWorkersMax, nil
}

func getMonth(arguments map[string]interface{}) string {
	var month string

	if arguments["--month"] != nil {
		month = arguments["--month"].(string)
	} else {
		month = time.Now().Format("Jan")
	}

	return month
}
