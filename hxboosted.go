package main

import (
	"fmt"

	"github.com/flamego/flamego"
	"github.com/flamego/template"
)

func HXBoosted(c flamego.Context, d template.Data) {
	boosted := c.Request().Request.Header.Get("HX-Boosted")
	d["Boosted"] = ""
	if boosted != "" {
		d["Boosted"] = "boosted"
		fmt.Println("Upgrading request to hx-boosted response.")
	}
}
