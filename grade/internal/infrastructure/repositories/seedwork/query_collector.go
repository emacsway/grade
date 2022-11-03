package seedwork

import (
	"errors"
	"strings"

	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type MultiQuerier interface {
	infrastructure.MutableQueryExecutor
	infrastructure.DbSessionExecutor
}

type QueryCollector struct {
	multiQueryMap map[string]MultiQuerier
}

func (c *QueryCollector) Exec(query string, args ...any) (infrastructure.Result, error) {
	if _, found := c.multiQueryMap[query]; !found {
		if strings.TrimSpace(query)[:6] == "INSERT" {
			c.multiQueryMap[query] = NewMultiInsertQuery()
		}
	}
	if multiQuery, found := c.multiQueryMap[query]; found {
		return multiQuery.Exec(query, args...)
	}
	return nil, errors.New("unknown SQL command")
}

func (c *QueryCollector) Execute(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	result := &DeferredResult{}
	var lastInsertId int64
	var rowsAffected int64
	for k := range c.multiQueryMap {
		r, err := c.multiQueryMap[k].Execute(s)
		if err != nil {
			return nil, err
		}
		rowsAffectedIncrement, err := r.RowsAffected()
		rowsAffected += rowsAffectedIncrement
		if err != nil {
			return nil, err
		}
		lastInsertId, err = r.LastInsertId()
		if err != nil {
			return nil, err
		}
		result.Resolve(rowsAffected, lastInsertId)
	}
	return result, nil
}
