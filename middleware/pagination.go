package middleware

import (
	"strconv"
	"strings"

	"github.com/LockBlock-dev/planes-tracker-api/constant"
	"github.com/LockBlock-dev/planes-tracker-api/context"
	"github.com/labstack/echo/v4"
	eMiddleware "github.com/labstack/echo/v4/middleware"
)

type PaginationConfig struct {
	Skipper eMiddleware.Skipper
}

var DefaultPaginationConfig = PaginationConfig{
	Skipper: eMiddleware.DefaultSkipper,
}

func Pagination() echo.MiddlewareFunc {
	c := DefaultPaginationConfig
	return PaginationWithConfig(c)
}

func PaginationWithConfig(config PaginationConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultPaginationConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			page, _ := strconv.Atoi(c.QueryParam("page"))
			if page <= 0 {
				page = constant.PAGINATION_DEFAULT_PAGE
			}

			pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
			switch {
			case pageSize > constant.PAGINATION_MAX_PAGE_SIZE:
				pageSize = constant.PAGINATION_MAX_PAGE_SIZE
			case pageSize <= 0:
				pageSize = constant.PAGINATION_DEFAULT_PAGE_SIZE
			}

			order := strings.ToUpper(c.QueryParam("order"))
			if order != constant.ORDERING_ASCENDING && order != constant.ORDERING_DESCENDING {
				order = constant.ORDERING_ASCENDING
			}

			sort := c.QueryParam("sort")

			filter := c.QueryParam("filter")

			paginationContext := context.NewPaginationContextWithoutDefaults(c, page, pageSize, order, sort, filter)

			return next(paginationContext)
		}
	}
}
