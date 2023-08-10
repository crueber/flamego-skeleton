package routes

import (
	"fmt"
	"net/http"

	"github.com/flamego/binding"
	"github.com/flamego/flamego"
	"github.com/flamego/template"
	"gorm.io/gorm"

	m "fgo-test/models"
)

func UserIndex(c flamego.Context, t template.Template, d template.Data, db *gorm.DB) {
	paging := m.SetPagination(c.QueryInt64("page"), c.QueryInt("perpage"), "/users")
	d["Users"], d["Pagination"] = m.PaginatedList[m.User](paging, db)
	t.HTML(http.StatusOK, "user/index")
}

func UserRead(c flamego.Context, t template.Template, d template.Data, db *gorm.DB) {
	d["User"] = m.Get[m.User](c.ParamInt("id"), db)
	t.HTML(http.StatusOK, "user/read")
}

func UserUpdate(c flamego.Context, t template.Template, d template.Data, userUpdates m.User, db *gorm.DB, errs binding.Errors) {
	if len(errs) > 0 {
		validationError(c, errs)
		return
	}
	updates := map[string]interface{}{
		"name":  userUpdates.Name,
		"email": userUpdates.Email,
	}
	user := Update[m.User](c.ParamInt("id"), updates, db)
	d["User"] = user
	if c.Request().Method == http.MethodPut {
		t.HTML(http.StatusAccepted, "user/partial")
	} else {
		c.Redirect(fmt.Sprintf("%s%d", "/user/", user.ID), http.StatusFound)
	}
}

func UserCreate(c flamego.Context, t template.Template, d template.Data, user m.User, db *gorm.DB, errs binding.Errors) {
	if len(errs) > 0 {
		validationError(c, errs)
		return
	}
	m.Create(&user, db)
	d["User"] = user
	t.HTML(http.StatusCreated, "user/partial")
}

func UserDelete(c flamego.Context, db *gorm.DB) {
	user := m.Get[m.User](c.ParamInt("id"), db)
	m.Delete(user, db)
	if c.Request().Method == http.MethodGet {
		c.Redirect("/users", http.StatusFound)
	} else {
		StatusOK(c)
	}
}
