package seedwork

import (
	"fmt"
	"regexp"

	"github.com/emacsway/grade/grade/internal/infrastructure"
)

var reInsert = regexp.MustCompile(`VALUES\s*(\((?:'(?:[^']|'')+'|[^)])+\))`)

func NewMultiInsertQuery() *MultiQuery {
	r := &MultiQuery{
		re:      reInsert,
		replace: "VALUES %s",
		concat:  ", ",
	}
	return r
}

type MultiQuery struct {
	sqlTemplate  string
	placeholders string
	params       [][]any
	results      []*DeferredResult
	re           *regexp.Regexp
	replace      string
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
	return fmt.Sprintf(q.sqlTemplate, bulkPlaceholders)
}

func (q *MultiQuery) flatParams() []any {
	var result []any
	for i := range q.params {
		result = append(result, q.params[i]...)
	}
	return result
}

func (q *MultiQuery) Exec(query string, args ...any) (infrastructure.Result, error) {
	q.placeholders = q.re.FindStringSubmatch(query)[1]
	q.sqlTemplate = q.re.ReplaceAllLiteralString(query, q.replace)
	q.params = append(q.params, args)
	result := &DeferredResult{}
	q.results = append(q.results, result)
	return result, nil
}

func (q MultiQuery) Execute(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	r, err := s.Exec(q.sql(), q.flatParams()...)
	if err != nil {
		return nil, err
	}
	// TODO: implement me.
	for i := range q.results {
		q.results[i].Resolve(0, 0)
	}
	return r, nil
}
