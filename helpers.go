package main

import (
	"html/template"
	"time"

	"github.com/dustin/go-humanize"
)

func DateOnly(t time.Time) string {
	return t.Format("01/02/2006")
}

func RelativeTime(t time.Time) string {
	return humanize.Time(t)
}

func Helpers() template.FuncMap {
	return template.FuncMap{
		"RelativeTime": RelativeTime,
		"DateOnly":     DateOnly,
	}
}
