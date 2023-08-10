package main

import (
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"

	m "fgo-test/models"
)

func Helpers() template.FuncMap {
	return template.FuncMap{
		"RelativeTime": RelativeTime,
		"DateOnly":     DateOnly,
		"Add":          Add,
		"BuildURI":     BuildURI,
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

func BuildURI(p m.Pagination) string {
	var uri strings.Builder
	uri.WriteString(p.URI)
	uri.WriteString("?")
	if p.NextPage != 0 {
		uri.WriteString("page=")
		uri.WriteString(strconv.FormatInt(p.NextPage, 10))
	}
	if p.PerPage != 5 && p.PerPage != 0 {
		uri.WriteString("&perpage=")
		uri.WriteString(strconv.Itoa(p.PerPage))
	}
	if p.Search != "" {
		uri.WriteString("&search=")
		uri.WriteString(p.Search)
	}
	return uri.String()
}
