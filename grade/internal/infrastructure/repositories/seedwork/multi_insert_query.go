package seedwork

import (
	"fmt"
	"regexp"

	"github.com/emacsway/grade/grade/internal/infrastructure"
)

var re = regexp.MustCompile(`VALUES\s*(\((?:'(?:[^']|'')+'|[^)])+\))`)

type MultiInsertQuery struct {
	sqlTemplate  string
	placeholders string
	params       [][]any
	result       *DeferredResult
}

func (q *MultiInsertQuery) sql() string {
	bulkPlaceholders := ""
	for i := 0; i < len(q.params); i++ {
		if i != 0 {
			bulkPlaceholders += ", "
		}
		bulkPlaceholders += q.placeholders
	}
	return fmt.Sprintf(q.sqlTemplate, bulkPlaceholders)
}

func (q *MultiInsertQuery) flatParams() []any {
	var result []any
	for i := range q.params {
		result = append(result, q.params[i][:]...)
	}
	return result
}

func (q *MultiInsertQuery) Exec(query string, args ...any) (infrastructure.Result, error) {
	q.placeholders = re.FindStringSubmatch(query)[1]
	q.sqlTemplate = re.ReplaceAllLiteralString(query, "VALUES %s")
	q.params = append(q.params, args)
	q.result = &DeferredResult{}
	return q.result, nil
}

func (q MultiInsertQuery) Execute(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	r, err := s.Exec(q.sql(), q.flatParams()...)
	if err != nil {
		return nil, err
	}
	// TODO: implement me.
	q.result.SetLastInsertId(0)
	q.result.SetRowsAffected(0)
	return r, nil
}
