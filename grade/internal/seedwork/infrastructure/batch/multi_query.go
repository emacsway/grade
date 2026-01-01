package batch

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/go-multierror"

	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/utils"
)

var reInsert = regexp.MustCompile(`VALUES\s*(\((?:'(?:[^']|'')*'|[^)])+\))`)

func NewMultiInsertQuery() *MultiQuery {
	r := &MultiQuery{
		MultiQueryBase{
			re:          reInsert,
			replacement: "VALUES %s",
			concat:      ", ",
		},
	}
	return r
}

func NewAutoincrementMultiInsertQuery() *AutoincrementMultiInsertQuery {
	r := &AutoincrementMultiInsertQuery{
		MultiQueryBase{
			re:          reInsert,
			replacement: "VALUES %s",
			concat:      ", ",
		},
	}
	return r
}

type MultiQueryBase struct {
	sqlTemplate  string
	placeholders string
	params       [][]any
	results      []*session.DeferredResultImp
	re           *regexp.Regexp
	replacement  string
	concat       string
}

func (q *MultiQueryBase) sql() string {
	bulkPlaceholders := ""
	for i := 0; i < len(q.params); i++ {
		if i != 0 {
			bulkPlaceholders += q.concat
		}
		bulkPlaceholders += q.placeholders
	}
	return utils.Rebind(fmt.Sprintf(q.sqlTemplate, bulkPlaceholders))
}

func (q *MultiQueryBase) flatParams() []any {
	var result []any
	for i := range q.params {
		result = append(result, q.params[i]...)
	}
	return result
}

func (q *MultiQueryBase) Exec(query string, args ...any) (session.DeferredResult, error) {
	query = utils.RebindReverse(query)
	q.placeholders = q.re.FindStringSubmatch(query)[1]
	q.sqlTemplate = q.re.ReplaceAllLiteralString(query, q.replacement)
	q.params = append(q.params, args)
	result := session.NewDeferredResult()
	q.results = append(q.results, result)
	return result, nil
}

type MultiQuery struct {
	MultiQueryBase
}

func (q MultiQuery) Evaluate(s session.DbSession) (session.Result, error) {
	var errs error
	r, err := s.Exec(q.sql(), q.flatParams()...)
	if err != nil {
		return nil, err
	}
	for i := range q.results {
		d := q.results[i]
		d.Resolve(0, 0)
		occurredErr := d.OccurredErr()
		if occurredErr != nil {
			errs = multierror.Append(errs, occurredErr)
		}
	}
	return r, errs
}

type AutoincrementMultiInsertQuery struct {
	MultiQueryBase
}

func (q AutoincrementMultiInsertQuery) Evaluate(s session.DbSession) (session.Result, error) {
	var id int64
	var errs error
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
		d := q.results[i]
		d.Resolve(id, 0)
		occurredErr := d.OccurredErr()
		if occurredErr != nil {
			errs = multierror.Append(errs, occurredErr)
		}
		i++
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return session.NewResult(0, int64(len(q.results))), errs
}
