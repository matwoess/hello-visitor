package main

import (
	"testing"
	"time"
)

func TestGetDateTimeString(t *testing.T) {
	year := 2022
	month := time.Month(2)
	day := 8
	hour := 0
	min := 9
	sec := 4
	nano := 0
	loc := time.Local
	someTime := time.Date(year, month, day, hour, min, sec, nano, loc)
	res := getDateTimeString(someTime)
	if res != "08.02.2022 00:09:04" {
		t.Errorf("getDateTimeString(someTime) = %s; wanted 08.02.2022 00:09:04", res)
	}
}
