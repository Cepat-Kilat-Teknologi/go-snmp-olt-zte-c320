package pagination

import (
	"net/http"
	"strconv"
)

var (
	DefaultPageSize = 10
	MaxPageSize     = 100
	PageVar         = "page"
	PageSizeVar     = "limit"
)

type Pages struct {
	Code      int32       `json:"code"`
	Status    string      `json:"status"`
	Page      int         `json:"page"`
	PageSize  int         `json:"limit"`
	PageCount int         `json:"page_count"`
	TotalRows int         `json:"total_rows"`
	Data      interface{} `json:"data"`
}

func New(page, pageSize, total int) *Pages {
	if page <= 0 {
		page = 0
	}
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + pageSize - 1) / pageSize
	}
	return &Pages{
		Code:      200,
		Status:    "OK",
		Page:      page,
		PageSize:  pageSize,
		TotalRows: total,
		PageCount: pageCount,
	}
}

func GetPaginationParametersFromRequest(r *http.Request) (pageIndex, pageSize int) {
	pageIndex = parseInt(r.URL.Query().Get(PageVar), 1)
	pageSize = parseInt(r.URL.Query().Get(PageSizeVar), DefaultPageSize)
	return pageIndex, pageSize
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
