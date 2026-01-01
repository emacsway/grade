package batch

import (
	"errors"

	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
)

func NewQueryCollector() *QueryCollector {
	r := &QueryCollector{
		multiQueryMap: make(map[string]MultiQuerier),
	}
	return r
}

type MultiQuerier interface {
	session.QueryEvaluator
	session.DeferredDbSessionExecutor
}

type QueryCollector struct {
	multiQueryMap map[string]MultiQuerier
}

func (c *QueryCollector) Exec(query string, args ...any) (session.DeferredResult, error) {
	if _, found := c.multiQueryMap[query]; !found {
		if session.IsAutoincrementInsertQuery(query) {
			c.multiQueryMap[query] = NewAutoincrementMultiInsertQuery()
		} else if session.IsInsertQuery(query) {
			c.multiQueryMap[query] = NewMultiInsertQuery()
		}
	}
	if multiQuery, found := c.multiQueryMap[query]; found {
		return multiQuery.Exec(query, args...)
	}
	return nil, errors.New("unknown SQL command")
}

func (c *QueryCollector) Evaluate(s session.DbSession) (session.Result, error) {
	var rowsAffected int64
	for len(c.multiQueryMap) > 0 {
		// Resolve N+1 query problem with auto-increment primary key.
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
	return session.NewResult(0, rowsAffected), nil
}
