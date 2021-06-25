package entity

import (
	"errors"
	"strings"
)

type (
	Pagination struct {
		Page  int `json:"page"`
		Row   int `json:"row"`
		Count int `json:"count"`
	}
)

var (
	ErrorDataTableNotValid = errors.New("DataTable request is not valid")
)

type (
	FilterDataTable struct {
		Sort       Sort
		Pagination PaginationFilter
	}
)

func (filter *FilterDataTable) Validate() {
	filter.Sort.Validate()
	filter.Pagination.Validate()
}

func (filter *FilterDataTable) IsValid() bool {
	if !filter.IsPaginated() {
		return true
	}

	if filter.Pagination.Limit == 0 {
		return false
	}

	if filter.Sort.Field == "" || filter.Sort.Direction == "" {
		return false
	}

	return true
}

func (filter *FilterDataTable) IsPaginated() bool {
	return !filter.Pagination.IsDisabled()
}

type (
	Sort struct {
		Field     string
		Direction string
	}
)

func (sort *Sort) Validate() {
	if !strings.EqualFold(sort.Direction, "desc") {
		sort.Direction = "asc"
	}
}

type (
	PaginationFilter struct {
		DisablePagination bool
		Page              int
		Limit             int
		Offset            int
	}
)

func (pagination *PaginationFilter) Validate() {
	if pagination.IsDisabled() {
		return
	}

	if !(pagination.Limit > 0) {
		pagination.Limit = 10
	}

	pagination.SetOffset()
}

func (pagination *PaginationFilter) SetOffset() {
	if pagination.IsDisabled() {
		return
	}

	if pagination.Page > 0 {
		pagination.Offset = (pagination.Page - 1) * pagination.Limit
	}
}

func (pagination *PaginationFilter) IsDisabled() bool {
	return pagination.DisablePagination
}
