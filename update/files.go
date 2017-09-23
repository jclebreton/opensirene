package main

import (
	"fmt"
	"math"
	"time"
)

//getFiles returns all the files to start from scratch
func getZipListFromScratch(monthName, url, dest string) (files []zipFile, err error) {

	//Start date
	firstDayOfMonth, err := getStartingDate(monthName)
	if err != nil {
		return nil, err
	}

	//Incremental files
	n := getNumOfDaysForMonth(firstDayOfMonth)
	day := firstDayOfMonth
	for i := 1; i <= n; i++ {
		dayName := day.Format("Monday")

		//No files the weekend
		if dayName != "Saturday" && dayName != "Sunday" {

			//No files in the future
			if time.Now().After(day) {
				file := getIncrementalFile(day, dest)
				files = append(files, file)
			}
		}
		day = firstDayOfMonth.AddDate(0, 0, i)
	}

	//Stock file
	file := getCompleteFile(firstDayOfMonth.AddDate(0, -1, 0).Add(time.Nanosecond), dest)
	files = append(files, file)

	return files, nil
}

//getNumOfDays returns the number of the days in the month for current date
func getNumOfDaysForMonth(d time.Time) int {
	t := time.Date(d.Year(), d.Month(), 32, 0, 0, 0, 0, time.UTC)
	return 32 - t.Day()
}

//getDayNumber returns the day's number in the year
func getDayNumber(d time.Time) int {
	year, _, _ := d.Date()
	first := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)

	return int(math.Ceil(d.Local().Sub(first).Hours() / 24))
}

//getStartingDate returns the corresponding time object for asked month
func getStartingDate(month string) (time.Time, error) {
	y, _, _ := time.Now().Date()
	layout := "2006-Jan-02"
	pattern := "%d-%s-01"

	//Month of current year
	date, err := time.Parse(layout, fmt.Sprintf(pattern, y, month))
	date = date.Add(time.Nanosecond)
	if err != nil {
		return time.Time{}, err
	}

	if date.Before(time.Now()) {
		return date, nil
	}

	//Month of last year
	date, err = time.Parse(layout, fmt.Sprintf(pattern, y-1, month))
	date = date.Add(time.Nanosecond)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

//getIncrementalFile returns the incremental file for the current date
func getIncrementalFile(d time.Time, dest string) zipFile {
	y, _, _ := d.Date()
	n := getDayNumber(d)
	return getFile(fmt.Sprintf("sirene_%d%03d_E_Q", y, n), "incremental", dest)
}

//getCompleteFile returns the stock file for the current date
func getCompleteFile(d time.Time, dest string) zipFile {
	y, m, _ := d.Date()
	return getFile(fmt.Sprintf("sirene_%d%02d_L_M", y, m), "complete", dest)
}

//getFile returns zip file with all meta data
func getFile(name, updateType, dest string) zipFile {
	file := zipFile{
		name:       name,
		updateType: updateType,
		filename:   fmt.Sprintf("%s.zip", name),
	}
	file.url = fmt.Sprintf("%s/%s", url, file.filename)
	file.path = fmt.Sprintf("%s/%s", dest, file.filename)
	return file
}
