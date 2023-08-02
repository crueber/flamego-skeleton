package routes

import (
  "fmt"
  "net/http"

  "github.com/flamego/flamego"
  "github.com/flamego/template"
  "github.com/flamego/binding"
  "gorm.io/gorm"

  m "fgo-test/models"
)

func UserIndex(c flamego.Context, t template.Template, d template.Data, db *gorm.DB) {
  d["Users"] = m.ListAllUsers(db)
  t.HTML(http.StatusOK, "user/index")
}

func UserRead(c flamego.Context, t template.Template, d template.Data, db *gorm.DB) {
  d["User"] = m.GetUser(db, c.ParamInt("id"))
  t.HTML(http.StatusOK, "user/read")
}

func UserUpdate(c flamego.Context, t template.Template, d template.Data, userUpdates m.User, db *gorm.DB, errs binding.Errors) {
  if len(errs) > 0 {
    validationError(c, errs)
    return
  }
  user := m.GetUser(db, c.ParamInt("id"))
  user.Name = userUpdates.Name
  user.Email = userUpdates.Email
  user.Save(db)
  d["User"] = user
  if c.Request().Method == http.MethodPut {
    t.HTML(http.StatusOK, "user/partial")
  } else {
    c.Redirect(fmt.Sprintf("%s%d", "/user/", user.ID), http.StatusFound)
  }
}

func UserCreate(c flamego.Context, t template.Template, d template.Data, user m.User, db *gorm.DB, errs binding.Errors) {
  if len(errs) > 0 {
    validationError(c, errs)
    return
  }
  user.Save(db)
  d["User"] = user
  t.HTML(http.StatusOK, "user/partial")
  // c.Redirect(fmt.Sprintf("%s%d", "/user/", user.ID), http.StatusFound)
}

func UserDelete(c flamego.Context, db *gorm.DB) {
  user := m.GetUser(db, c.ParamInt("id"))
  user.Delete(db)
  if c.Request().Method == http.MethodGet {
    c.Redirect("/users", http.StatusFound)
  } else {
    w := c.ResponseWriter()
    w.WriteHeader(http.StatusOK)
    w.Write([]byte{})
  }
}

