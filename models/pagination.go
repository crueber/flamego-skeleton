package models

type Pagination struct {
	TotalPages  int64
	CurrentPage int64
	PerPage     int
	URI         string
	PrevPage    int64
	NextPage    int64
}

func (p *Pagination) Offset() int {
	return int((p.CurrentPage - 1) * int64(p.PerPage))
}

func (p *Pagination) SetTotal(count int64) {
	p.TotalPages = (count / int64(p.PerPage)) + 1

	if p.CurrentPage > 1 {
		p.PrevPage = p.CurrentPage - 1
	} else {
		p.PrevPage = 0
	}

	if p.CurrentPage < p.TotalPages {
		p.NextPage = p.CurrentPage + 1
	} else {
		p.NextPage = 0
	}
}

func SetPagination(page int64, per int, uri string) Pagination {
	if page == 0 {
		page = 1
	}
	if per == 0 {
		per = 5
	}

	return Pagination{0, page, per, uri, 0, 0}
}
