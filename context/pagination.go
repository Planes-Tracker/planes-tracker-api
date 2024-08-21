package context

import (
	"github.com/LockBlock-dev/planes-tracker-api/constant"
	"github.com/labstack/echo/v4"
)

type PaginationContext struct {
	echo.Context
	page     int
	pageSize int
	order    string
	sort     string
	filter   string
}

func (c *PaginationContext) Page() int {
	return c.page
}

func (c *PaginationContext) SetPage(page int) *PaginationContext {
	c.page = page

	return c
}

func (c *PaginationContext) PageSize() int {
	return c.pageSize
}

func (c *PaginationContext) SetPageSize(pageSize int) *PaginationContext {
	c.pageSize = pageSize

	return c
}

func (c *PaginationContext) Order() string {
	return c.order
}

func (c *PaginationContext) SetOrder(order string) *PaginationContext {
	c.order = order

	return c
}

func (c *PaginationContext) Sort() string {
	return c.sort
}

func (c *PaginationContext) SetSort(sort string) *PaginationContext {
	c.sort = sort

	return c
}

func (c *PaginationContext) Filter() string {
	return c.filter
}

func (c *PaginationContext) SetFilter(filter string) *PaginationContext {
	c.filter = filter

	return c
}

func NewPaginationContext(c echo.Context) *PaginationContext {
	return &PaginationContext{
		c,
		constant.PAGINATION_DEFAULT_PAGE,
		constant.PAGINATION_DEFAULT_PAGE_SIZE,
		constant.ORDERING_ASCENDING,
		"",
		"",
	}
}

func NewPaginationContextWithoutDefaults(c echo.Context, page int, pageSize int, order string, sort string, filter string) *PaginationContext {
	return &PaginationContext{c, page, pageSize, order, sort, filter}
}
