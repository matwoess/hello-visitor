package main

import (
	"testing"
	"time"
)

func TestGetDateTimeString(t *testing.T) {
	someTime := time.Date(2022, 12, 28, 18, 19, 45, 2, time.Local)
	res := getDateTimeString(someTime)
	if res != "28.12.2022 18:19:45" {
		t.Errorf("getDateTimeString(someTime) = %s; wanted 28.12.2022 18:19:45", res)
	}
}
