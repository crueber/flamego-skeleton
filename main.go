package main

import (
  "github.com/flamego/flamego"
  "github.com/flamego/template"
  "github.com/flamego/binding"
  "github.com/flamego/session"

  "fgo-test/templates"
  m "fgo-test/models"
  r "fgo-test/routes"
)

func main() {
  fs, _ := template.EmbedFS(templates.Templates, ".", []string{".tmpl"})
  db := OpenDatabase()
  f := flamego.New()

  f.Map(db)
  f.Use(flamego.Logger())
  f.Use(flamego.Recovery())
  f.Use(flamego.Static(flamego.StaticOptions{Directory: "public"}))
  f.Use(template.Templater(template.Options{ FileSystem: fs, }))
  f.Use(session.Sessioner())

  f.Post("/user", binding.Form(m.User{}), r.UserCreate)
  f.Post("/user/{id}", binding.Form(m.User{}), r.UserUpdate)
  f.Get("/users", r.UserIndex)
  f.Get("/user/{id}", r.UserRead)
  f.Get("/user/{id}/delete", r.UserDelete)

  f.Run()
}

