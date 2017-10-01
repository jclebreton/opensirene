package download_extract

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getEasterDay(t *testing.T) {
	date := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	easterDay := getEasterDay(date)
	assert.Equal(t, 2017, easterDay.Year())
	assert.Equal(t, time.Month(4), easterDay.Month())
	assert.Equal(t, 16, easterDay.Day())
}

func Test_isWorkingDay(t *testing.T) {
	year := time.Now().Year()

	nouvelAn := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	feteDuTravail := time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC)
	dDay := time.Date(year, 5, 8, 0, 0, 0, 0, time.UTC)
	ascension := time.Date(year, 5, 25, 0, 0, 0, 0, time.UTC)
	pentecote := time.Date(year, 6, 5, 0, 0, 0, 0, time.UTC)
	revolution := time.Date(year, 7, 14, 0, 0, 0, 0, time.UTC)
	assomption := time.Date(year, 8, 15, 0, 0, 0, 0, time.UTC)
	toussaint := time.Date(year, 11, 1, 0, 0, 0, 0, time.UTC)
	armistice := time.Date(year, 11, 11, 0, 0, 0, 0, time.UTC)
	noel := time.Date(year, 12, 25, 0, 0, 0, 0, time.UTC)
	lundiDePaques := getEasterDay(time.Now()).AddDate(0, 0, 1)
	sunday := getEasterDay(time.Now())
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
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 28, getNumOfDaysForMonth(time.Date(2017, 2, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 3, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 30, getNumOfDaysForMonth(time.Date(2017, 4, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 5, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 30, getNumOfDaysForMonth(time.Date(2017, 6, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 7, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 8, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 30, getNumOfDaysForMonth(time.Date(2017, 9, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 30, getNumOfDaysForMonth(time.Date(2017, 11, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, 31, getNumOfDaysForMonth(time.Date(2017, 12, 1, 0, 0, 0, 0, time.UTC)))
}

func Test_getDayNumber(t *testing.T) {
	day := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 365; i++ {
		assert.Equal(t, i+1, getDayNumber(day.AddDate(0, 0, i)))
	}
}

func Test_getStartingDate(t *testing.T) {
	currentMonth := time.Now()
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
