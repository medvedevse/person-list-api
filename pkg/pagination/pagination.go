package pagination

import "gorm.io/gorm"

type PaginationFilters struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

func InitPagination(limit, page int) *PaginationFilters {
	return &PaginationFilters{Limit: limit, Page: page}
}

func (p *PaginationFilters) GetPaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit
	return db.Offset(offset).Limit(p.Limit)
}
