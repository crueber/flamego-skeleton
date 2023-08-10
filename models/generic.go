package models

import "gorm.io/gorm"

type Model interface {
	User
}

func ListAll[M Model](db *gorm.DB) []M {
	var hydrated []M
	db.Find(&hydrated)
	return hydrated
}

func Get[M Model](id int, db *gorm.DB) M {
	var hydrated M
	db.First(&hydrated, id)
	return hydrated
}

func Create[M Model](model *M, db *gorm.DB) error {
	result := db.Create(model)
	return result.Error
}

func Delete[M Model](model M, db *gorm.DB) error {
	result := db.Delete(&model)
	return result.Error
}
