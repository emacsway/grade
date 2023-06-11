package seedwork

import (
	"errors"
	"strings"

	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type MultiQuerier interface {
	infrastructure.MutableQueryEvaluator
	infrastructure.DeferredDbSessionExecutor
}

type QueryCollector struct {
	multiQueryMap map[string]MultiQuerier
}

func (c *QueryCollector) Exec(query string, args ...any) (infrastructure.DeferredResult, error) {
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

func (c *QueryCollector) Evaluate(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	var rowsAffected int64
	for len(c.multiQueryMap) > 0 {
		// Nested queries have got the lastInsertId and can be handled for now
		currentMultiQueryMap := c.multiQueryMap
		c.multiQueryMap = make(map[string]MultiQuerier)
		for k := range currentMultiQueryMap {
			r, err := currentMultiQueryMap[k].Evaluate(s)
			if err != nil {
				return nil, err
			}
			rowsAffectedIncrement, err := r.RowsAffected()
			if err == nil {
				rowsAffected += rowsAffectedIncrement
			}
		}
	}
	return infrastructure.RowsAffected(rowsAffected), nil
}
