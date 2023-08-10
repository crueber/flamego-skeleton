package routes

import (
	"fmt"
	"net/http"

	"github.com/flamego/binding"
	"github.com/flamego/flamego"
	"github.com/flamego/validator"
	"gorm.io/gorm"

	m "fgo-test/models"
)

func StatusOK(c flamego.Context) {
	w := c.ResponseWriter()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}

func Update[M m.Model](id int, updates map[string]interface{}, db *gorm.DB) M {
	db.Model(&M{ID: uint(id)}).Updates(updates)
	return m.Get[M](id, db)
}

func validationError(c flamego.Context, errs binding.Errors) {
	var err error
	switch errs[0].Category {
	case binding.ErrorCategoryValidation:
		err = errs[0].Err.(validator.ValidationErrors)[0]
	default:
		err = errs[0].Err
	}

	w := c.ResponseWriter()
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf("Error: %v", err)))
}
