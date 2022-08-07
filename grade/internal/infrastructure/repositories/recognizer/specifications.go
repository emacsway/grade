package recognizer

import (
	"database/sql/driver"
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/recognizer"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	s "github.com/emacsway/grade/grade/internal/infrastructure/specification"
)

type RecognizerCanCompleteEndorsementSpecification struct {
	recognizer.RecognizerCanCompleteEndorsementSpecification
}

func (r *RecognizerCanCompleteEndorsementSpecification) Compile() (sql string, params []driver.Valuer, err error) {
	v := s.NewPostgresqlVisitor(Context{})
	err = r.Expression().Accept(v)
	if err != nil {
		return "", []driver.Valuer{}, err
	}
	return v.Result()
}

type Context struct {
}

func (c Context) NameByPath(path ...string) (string, error) {
	switch path[0] {
	case "recognizer":
		return c.recognizerPath("recognizer", path[1:]...)
	default:
		return "", fmt.Errorf("can't get object \"%s\"", path[0])
	}
}

func (c Context) recognizerPath(prefix string, path ...string) (string, error) {
	switch path[0] {
	case "availableEndorsementCount":
		return prefix + ".available_endorsement_count", nil
	case "pendingEndorsementCount":
		return prefix + ".pending_endorsement_count", nil
	default:
		return "", fmt.Errorf("can't get field \"%s\"", path[0])
	}
}

func (c Context) Extract(val any) (driver.Valuer, error) {
	switch valTyped := val.(type) {
	case recognizer.EndorsementCount:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return ex, nil
	default:
		return nil, fmt.Errorf("can't export \"%#v\"", val)
	}
}
