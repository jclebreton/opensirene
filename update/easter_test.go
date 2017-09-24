package main

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
