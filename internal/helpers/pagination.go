package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationParams struct {
	Page     int
	PageSize int
}

type MetaData struct {
	TotalRecords int64 `json:"total_records"`
	Page         int   `json:"page"`
	Offset       int   `json:"offset"`
	Limit        int   `json:"limit"`
	TotalPages   int   `json:"total_pages"`
}

func GetPaginationParams(c *gin.Context) PaginationParams {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	defaultPage, defaultPageSize := 1, 10

	if page < 1 {
		page = defaultPage
	}

	if pageSize < 1 {
		pageSize = defaultPageSize
	}

	return PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}
}

func (p *PaginationParams) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (m *MetaData) CalculateTotalPage() int {
	return (int(m.TotalRecords) + m.Limit - 1) / m.Limit
}
