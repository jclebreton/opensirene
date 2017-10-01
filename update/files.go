package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//getFiles returns all the files to start from scratch
func getScratchZipList(monthName, url, dest string) (files []zipFile, err error) {

	//Start date
	firstDayOfMonth, err := getStartingDate(monthName)
	if err != nil {
		return nil, err
	}

	//Incremental files
	n := getNumOfDaysForMonth(firstDayOfMonth)
	day := firstDayOfMonth
	for i := 1; i <= n; i++ {
		if isWorkingDay(day) && time.Now().After(day) {
			file := getIncrementalFile(day, dest, url)
			files = append(files, file)
		}
		day = firstDayOfMonth.AddDate(0, 0, i)
	}

	//Stock file
	file := getCompleteFile(firstDayOfMonth.AddDate(0, -1, 0).Add(time.Nanosecond), dest, url)
	files = append(files, file)

	return files, nil
}

//getFiles returns all the files to start from scratch
func getDailyZipList(url, dest string) (files []zipFile, err error) {

	day := time.Now().AddDate(0, 0, -1)
	if !isWorkingDay(day) {
		return nil, errors.New("Not a working day")
	}

	file := getIncrementalFile(day, dest, url)
	if file.remoteExist() {
		return append(files, file), nil
	}

	return nil, errors.New("No daily file")
}

func getZipFileNames(files []zipFile) string {
	var l []string
	for _, f := range files {
		l = append(l, f.filename)
	}
	return strings.Join(l, ", ")
}

func getCsvFileNames(files []csvFile) string {
	var l []string
	for _, f := range files {
		l = append(l, f.filename)
	}
	return strings.Join(l, ", ")
}

//getIncrementalFile returns the incremental file for the current date
func getIncrementalFile(d time.Time, dest, url string) zipFile {
	y, _, _ := d.Date()
	n := getDayNumber(d)
	return getFile(fmt.Sprintf("sirene_%d%03d_E_Q", y, n), "incremental", dest, url)
}

//getCompleteFile returns the stock file for the current date
func getCompleteFile(d time.Time, dest, url string) zipFile {
	y, m, _ := d.Date()
	return getFile(fmt.Sprintf("sirene_%d%02d_L_M", y, m), "complete", dest, url)
}

//getFile returns zip file with all meta data
func getFile(name, updateType, dest, url string) zipFile {
	file := zipFile{
		name:       name,
		updateType: updateType,
		filename:   fmt.Sprintf("%s.zip", name),
	}
	file.url = fmt.Sprintf("%s/%s", url, file.filename)
	file.path = fmt.Sprintf("%s/%s", dest, file.filename)
	return file
}
