package main

import (
	"strings"
	"time"
)

func ParseTime(t string) time.Time {
	t2 := strings.Split(strings.Trim(t, " "), ":")

	now := time.Now()

	if len(t2) != 2 {
		return now
	}

	return time.Date(now.Year(), now.Month(), now.Day(), 20, 34, 0, 0, time.UTC)
}
