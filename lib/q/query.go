package q

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sunweiwe/paddle/core/common"
)

type KeyWords = map[string]interface{}

type Query struct {
	Keywords KeyWords

	Sorts []*Sort

	PageNumber int

	PageSize int

	WithoutPagination bool
}

func (q *Query) Limit() int {
	if q.WithoutPagination {
		return common.DefaultPageNumber
	}

	if q.PageSize < 1 {
		q.PageSize = common.DefaultPageSize
	}
	return q.PageSize
}

func (q *Query) Offset() int {
	if q.WithoutPagination {
		return 0
	}
	if q.PageSize < 1 {
		q.PageSize = common.DefaultPageSize
	}
	if q.PageNumber < 1 {
		q.PageNumber = common.DefaultPageNumber
	}
	return (q.PageNumber - 1) * q.PageSize
}

type Sort struct {
	Key  string
	DESC bool
}

func New(kw KeyWords) *Query {
	return &Query{Keywords: kw}
}

func (q *Query) WithPagination(c *gin.Context) *Query {
	pageNumber, _ := strconv.Atoi(c.Query(common.PageNumber))
	pageSize, _ := strconv.Atoi(c.Query(common.PageSize))

	if pageNumber < 1 {
		pageNumber = common.DefaultPageNumber
	}

	if pageSize < 0 {
		pageSize = common.DefaultPageSize
	}

	q.PageNumber = pageNumber
	q.PageSize = pageSize
	return q
}
