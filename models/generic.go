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

func PaginatedList[M Model](paging Pagination, db *gorm.DB) ([]M, Pagination) {
	var hydrated []M
	var count int64
	db.Limit(paging.PerPage).Offset(int((paging.CurrentPage - 1) * int64(paging.PerPage))).Find(&hydrated)
	db.Model(&M{}).Count(&count)
	paging.TotalPages = (count / int64(paging.PerPage)) + 1
	return hydrated, paging
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
