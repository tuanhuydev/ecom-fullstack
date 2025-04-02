package dto

type PaginationQueryDTO struct {
	Page      int    `form:"page"`
	PageSize  int    `form:"pageSize"`
	SortBy    string `form:"sortBy"`
	SortOrder string `form:"sortOrder"`
}
