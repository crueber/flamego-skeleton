package models

type Pagination struct {
	TotalPages  int64
	CurrentPage int64
	PerPage     int
}

func SetPagination(page int64, per int) Pagination {
	if page == 0 {
		page = 1
	}
	if per == 0 {
		per = 5
	}

	return Pagination{0, page, per}
}
