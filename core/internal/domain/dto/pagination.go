package dto

type Pagination struct {
	OrderBy string `form:"order"`
	SortBy  string `form:"sort"`
	Page    int    `form:"page"`
}
