package main

import (
	t "html/template"

	"github.com/flamego/binding"
	"github.com/flamego/flamego"
	"github.com/flamego/session"
	"github.com/flamego/template"

	m "fgo-test/models"
	r "fgo-test/routes"
	"fgo-test/templates"
)

func main() {
	fs, _ := template.EmbedFS(templates.Templates, ".", []string{".tmpl"})
	db := OpenDatabase()
	f := flamego.New()

	f.Map(db)
	f.Use(flamego.Logger())
	f.Use(flamego.Recovery())
	f.Use(flamego.Static(flamego.StaticOptions{Directory: "public"}))
	f.Use(template.Templater(template.Options{FuncMaps: []t.FuncMap{Helpers()}, FileSystem: fs}))
	f.Use(flamego.Renderer(flamego.RenderOptions{JSONIndent: "  "}))
	f.Use(HXBoosted)
	f.Use(session.Sessioner())

	f.Get("/", func(c flamego.Context) { c.Redirect("/users", 301) })
	f.Post("/user", binding.Form(m.User{}), r.UserCreate)
	f.Put("/user/{id}", binding.Form(m.User{}), r.UserUpdate)
	f.Post("/user/{id}", binding.Form(m.User{}), r.UserUpdate)
	f.Get("/users", r.UserIndex)
	f.Get("/users.partial", r.UserIndexPartial)
	f.Get("/user/{id}", r.UserRead)
	f.Delete("/user/{id}", r.UserDelete)
	f.Get("/user/{id}/delete", r.UserDelete)

	f.Run()
}
