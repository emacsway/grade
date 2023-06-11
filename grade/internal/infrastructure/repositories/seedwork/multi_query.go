package seedwork

import (
	"fmt"
	"regexp"

	"github.com/emacsway/grade/grade/internal/infrastructure"
)

var reInsert = regexp.MustCompile(`VALUES\s*(\((?:'(?:[^']|'')*'|[^)])+\))`)

func NewMultiInsertQuery() *MultiQuery {
	r := &MultiQuery{
		re:          reInsert,
		replacement: "VALUES %s",
		concat:      ", ",
	}
	return r
}

type MultiQuery struct {
	sqlTemplate  string
	placeholders string
	params       [][]any
	results      []*DeferredResult
	re           *regexp.Regexp
	replacement  string
	concat       string
}

func (q *MultiQuery) sql() string {
	bulkPlaceholders := ""
	for i := 0; i < len(q.params); i++ {
		if i != 0 {
			bulkPlaceholders += q.concat
		}
		bulkPlaceholders += q.placeholders
	}
	return Rebind(fmt.Sprintf(q.sqlTemplate, bulkPlaceholders))
}

func (q *MultiQuery) flatParams() []any {
	var result []any
	for i := range q.params {
		result = append(result, q.params[i]...)
	}
	return result
}

func (q *MultiQuery) Exec(query string, args ...any) (infrastructure.DeferredResult, error) {
	query = RebindReverse(query)
	q.placeholders = q.re.FindStringSubmatch(query)[1]
	q.sqlTemplate = q.re.ReplaceAllLiteralString(query, q.replacement)
	q.params = append(q.params, args)
	result := &DeferredResult{}
	q.results = append(q.results, result)
	return result, nil
}

func (q MultiQuery) Evaluate(s infrastructure.DbSession) (infrastructure.Result, error) {
	var id int64
	rows, err := s.Query(q.sql(), q.flatParams()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		q.results[i].Resolve(id, 0)
		i++
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return infrastructure.RowsAffected(len(q.results)), nil
}
