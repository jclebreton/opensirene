package download_extract

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getEasterDay(t *testing.T) {
	_ = SetParisLocation()
	date := time.Date(2017, 1, 1, 0, 0, 0, 1, location)
	easterDay := getEasterDay(date)
	assert.Equal(t, 2017, easterDay.Year())
	assert.Equal(t, time.Month(4), easterDay.Month())
	assert.Equal(t, 16, easterDay.Day())
}

func Test_isWorkingDay(t *testing.T) {
	_ = SetParisLocation()
	year := time.Now().In(location).Year()

	nouvelAn := time.Date(year, 1, 1, 0, 0, 0, 1, location)
	feteDuTravail := time.Date(year, 5, 1, 0, 0, 0, 1, location)
	dDay := time.Date(year, 5, 8, 0, 0, 0, 1, location)
	ascension := time.Date(year, 5, 25, 0, 0, 0, 1, location)
	pentecote := time.Date(year, 6, 5, 0, 0, 0, 1, location)
	revolution := time.Date(year, 7, 14, 0, 0, 0, 1, location)
	assomption := time.Date(year, 8, 15, 0, 0, 0, 1, location)
	toussaint := time.Date(year, 11, 1, 0, 0, 0, 1, location)
	armistice := time.Date(year, 11, 11, 0, 0, 0, 1, location)
	noel := time.Date(year, 12, 25, 0, 0, 0, 1, location)
	lundiDePaques := getEasterDay(time.Now().In(location)).AddDate(0, 0, 1)
	sunday := getEasterDay(time.Now().In(location))
	saturday := sunday.AddDate(0, 0, -1)
	monday := lundiDePaques.AddDate(0, 0, 7)
	thuesday := lundiDePaques.AddDate(0, 0, 1)

	assert.False(t, isWorkingDay(nouvelAn))
	assert.False(t, isWorkingDay(feteDuTravail))
	assert.False(t, isWorkingDay(dDay))
	assert.False(t, isWorkingDay(ascension))
	assert.False(t, isWorkingDay(pentecote))
	assert.False(t, isWorkingDay(revolution))
	assert.False(t, isWorkingDay(assomption))
	assert.False(t, isWorkingDay(toussaint))
	assert.False(t, isWorkingDay(armistice))
	assert.False(t, isWorkingDay(noel))
	assert.False(t, isWorkingDay(lundiDePaques))
	assert.False(t, isWorkingDay(sunday))
	assert.False(t, isWorkingDay(saturday))
	assert.True(t, isWorkingDay(thuesday))
	assert.True(t, isWorkingDay(monday))
}

func Test_getNumOfDaysForMonth(t *testing.T) {
	_ = SetParisLocation()
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 1, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 28, getNumOfDaysForMonth(time.Date(2017, 2, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 3, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 30, getNumOfDaysForMonth(time.Date(2017, 4, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 5, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 30, getNumOfDaysForMonth(time.Date(2017, 6, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 7, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 8, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 30, getNumOfDaysForMonth(time.Date(2017, 9, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 10, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 30, getNumOfDaysForMonth(time.Date(2017, 11, 1, 0, 0, 0, 1, location)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 12, 1, 0, 0, 0, 1, location)))
}

func Test_getDayNumber(t *testing.T) {
	_ = SetParisLocation()
	day := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 1, location)
	for i := 1; i <= 365; i++ {
		assert.Equal(t, i, getDayNumber(day.AddDate(0, 0, i-1)))
	}
}

func Test_getStartingDate(t *testing.T) {
	_ = SetParisLocation()
	currentMonth := time.Now().In(location)
	previousMonth := currentMonth.AddDate(0, -1, 0)
	nextMonth := currentMonth.AddDate(0, 1, 0)

	date, err := getStartingDate(nextMonth.Format("Jan"))
	assert.NoError(t, err)
	assert.Equal(t, currentMonth.Year()-1, date.Year())

	date, err = getStartingDate(previousMonth.Format("Jan"))
	assert.NoError(t, err)
	assert.Equal(t, previousMonth.Year(), date.Year())

	date, err = getStartingDate(currentMonth.Format("Jan"))
	assert.NoError(t, err)
	assert.Equal(t, currentMonth.Year(), date.Year())
}
