package domain

import "time"

type History struct {
	ID        int32
	Datetime  time.Time
	Filename  string
	IsSuccess bool
	Err       string
}