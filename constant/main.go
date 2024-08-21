package constant

import (
	"time"
)

const PAGINATION_DEFAULT_PAGE = 1
const PAGINATION_DEFAULT_PAGE_SIZE = 10
const PAGINATION_MAX_PAGE_SIZE = 100

const RESPONSE_DEFAULT_TIMEOUT = 10 * time.Second

const (
	ORDERING_ASCENDING  string = "ASC"
	ORDERING_DESCENDING string = "DESC"
)

const (
	FILTERS_MILITARY string = "MILITARY"
)
