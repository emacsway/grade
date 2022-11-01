package seedwork

import (
	"regexp"

	"github.com/emacsway/grade/grade/internal/infrastructure"
)

var re = regexp.MustCompile(`VALUES\s*(\((?:'(?:[^']|'')+'|[^)])+\))`)

type SingleInsertQuery interface {
	Execute(s infrastructure.DbSessionExecutor) (infrastructure.Result, error)
}

type MultiInsertQuery struct {
	sql          string
	placeholders string
	params       [][]any
}

func (mq *MultiInsertQuery) Add(sq SingleInsertQuery) {
	sq.Execute(mq)
}

func (mq *MultiInsertQuery) Exec(query string, args ...any) (infrastructure.Result, error) {
	mq.placeholders = re.FindStringSubmatch(query)[1]
	mq.sql = re.ReplaceAllLiteralString(query, "VALUES %s")
	mq.params = append(mq.params, args)
	return Result{}, nil
}

type Result struct {
	// TODO: make me a lazy object with real values
}

func (r Result) LastInsertId() (int64, error) {
	return 0, nil
}
func (r Result) RowsAffected() (int64, error) {
	return 0, nil
}
