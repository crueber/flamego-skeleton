package main

import (
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

func Helpers() template.FuncMap {
	return template.FuncMap{
		"RelativeTime":   RelativeTime,
		"DateOnly":       DateOnly,
		"Add":            Add,
		"BuildPagingURI": BuildPagingURI,
	}
}

func DateOnly(t time.Time) string {
	return t.Format("01/02/2006")
}

func RelativeTime(t time.Time) string {
	return humanize.Time(t)
}

func Add(n int64, a int64) int64 {
	return n + a
}

func BuildPagingURI(baseuri string, page int64, perpage int, search string) string {
	var uri strings.Builder
	uri.WriteString(baseuri)
	uri.WriteString("?")
	if page != 0 {
		uri.WriteString("page=")
		uri.WriteString(strconv.FormatInt(page, 10))
	}
	if perpage != 5 && perpage != 0 {
		uri.WriteString("&perpage=")
		uri.WriteString(strconv.Itoa(perpage))
	}
	if search != "" {
		uri.WriteString("&search=")
		uri.WriteString(search)
	}
	return uri.String()
}
