package main

import (
  "fmt"

  "github.com/flamego/flamego"
  "github.com/flamego/template"
)

func HXBoosted(c flamego.Context, d template.Data) {
  boosted := c.Request().Request.Header.Get("HX-Boosted")
  fmt.Println("%s", c.Request().Request.Header.Get("HX-Boosted"))
  if boosted != "" {
    d["Boosted"] = true
  }
}

