package download_extract

import (
	"fmt"
	"math"
	"time"
)

var location *time.Location

//getEasterDay returns the date of the easter day
func getEasterDay(date time.Time) time.Time {
	var a, b, c, d, e, r int
	year := date.Year()
	a = year % 19
	if year >= 1583 {
		var f, g, h, i, k, l, m int
		b = year / 100
		c = year % 100
		d = b / 4
		e = b % 4
		f = (b + 8) / 25
		g = (b - f + 1) / 3
		h = (19*a + b - d - g + 15) % 30
		i = c / 4
		k = c % 4
		l = (32 + 2*e + 2*i - h - k) % 7
		m = (a + 11*h + 22*l) / 451
		r = 22 + h + l - 7*m
	} else {
		b = year % 7
		c = year % 4
		d = (19*a + 15) % 30
		e = (2*c + 4*b - d + 34) % 7
		r = 22 + d + e
	}

	return time.Date(year, time.March, r, 0, 0, 0, 1, location)
}

//isWorkingDay returns true if current day is a working day
func isWorkingDay(day time.Time) bool {
	//Weekend
	dayName := day.Format("Monday")
	if dayName == "Saturday" || dayName == "Sunday" {
		return false
	}

	//Public holidays or the day after a holiday
	if isPublicDay(day) {
		return false
	}

	return true
}

//Public holidays
func isPublicDay(day time.Time) bool {
	easter := getEasterDay(day)
	return (day.Month() == 1 && day.Day() == 1) || //Jour de l'an
		(day.Month() == easter.Month() && day.Day() == easter.Day()+1) || //Lundi de Paques (avril)
		(day.Month() == 5 && day.Day() == 1) || //Fete du travail
		(day.Month() == 5 && day.Day() == 8) || //Liberation
		(day.Month() == 5 && day.Day() == 25) || //Ascension
		(day.Month() == 6 && day.Day() == 5) || //Pentecote
		(day.Month() == 7 && day.Day() == 14) || //Revolution
		(day.Month() == 8 && day.Day() == 15) || //Assomption
		(day.Month() == 11 && day.Day() == 1) || //Toussaint
		(day.Month() == 11 && day.Day() == 11) || //Armistice
		(day.Month() == 12 && day.Day() == 25) //Noel
}

//getNumOfDays returns the number of the days in the month for current date
func getNumOfDaysForMonth(d time.Time) int {
	t := time.Date(d.Year(), d.Month(), 32, 0, 0, 0, 1, location)
	return 32 - t.Day()
}

//getDayNumber returns the day's number in the year
func getDayNumber(d time.Time) int {
	year, _, _ := d.Date()
	first := time.Date(year, time.January, 1, 0, 0, 0, 1, d.Location())

	return int(math.Ceil(d.Sub(first).Hours()/24)) + 1
}

//getStartingDate returns the corresponding time object for asked month
func getStartingDate(month string) (time.Time, error) {
	y, _, _ := time.Now().In(location).Date()
	layout := "2006-Jan-02"
	pattern := "%d-%s-01"

	//Month of current year
	date, err := time.Parse(layout, fmt.Sprintf(pattern, y, month))
	date = date.Add(time.Nanosecond)
	if err != nil {
		return time.Time{}, err
	}

	if date.Before(time.Now().In(location)) {
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

func SetParisLocation() error {
	l, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		return err
	}
	location = l
	return nil
}
