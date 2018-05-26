package main

import (
	"strconv"
	"strings"
	"time"
)

func ParseTime(t string) time.Time {
	t2 := strings.Split(strings.Trim(t, " "), ":")

	now := time.Now()

	if len(t2) != 2 {
		return now
	}

	h, _ := strconv.ParseInt(t2[0], 10, 32)
	m, _ := strconv.ParseInt(t2[1], 10, 32)
	return time.Date(now.Year(), now.Month(), now.Day(), int(h), int(m), 0, 0, time.UTC)
}
