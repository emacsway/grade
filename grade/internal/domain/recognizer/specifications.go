package recognizer

import (
	"fmt"

	s "github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
)

type RecognizerCriteria struct{}

func (r RecognizerCriteria) availableEndorsementCount() s.FieldNode {
	return s.Field(r.obj(), "availableEndorsementCount")
}

func (r RecognizerCriteria) pendingEndorsementCount() s.FieldNode {
	return s.Field(r.obj(), "pendingEndorsementCount")
}

func (r RecognizerCriteria) obj() s.ObjectNode {
	return s.Object("recognizer")
}

var recognizer = RecognizerCriteria{}

type RecognizerCanCompleteEndorsementSpecification struct {
}

func (r RecognizerCanCompleteEndorsementSpecification) Expression() s.Visitable {
	return s.And(
		s.NotEqual(
			recognizer.availableEndorsementCount(),
			s.Value(EndorsementCount(0)),
		),
		s.NotEqual(
			recognizer.pendingEndorsementCount(),
			s.Value(EndorsementCount(0)),
		),
		s.GreaterThanEqual(
			recognizer.availableEndorsementCount(),
			recognizer.pendingEndorsementCount(),
		),
	)
}

func (r RecognizerCanCompleteEndorsementSpecification) IsSatisfiedBy(obj Recognizer) (bool, error) {
	v := s.NewEvaluateVisitor(Context{
		recognizer: obj,
	})
	err := r.Expression().Accept(v)
	if err != nil {
		return false, err
	}
	return v.Result()
}

type Context struct {
	recognizer Recognizer
}

func (c Context) ValuesByPath(path ...string) ([]any, error) {
	switch path[0] {
	case "recognizer":
		return c.recognizerPath(c.recognizer, path[1:]...)
	default:
		return []any{}, fmt.Errorf("can't get object \"%s\"", path[0])
	}
}

func (c Context) recognizerPath(obj Recognizer, path ...string) ([]any, error) {
	switch path[0] {
	case "availableEndorsementCount":
		return []any{obj.availableEndorsementCount}, nil
	case "pendingEndorsementCount":
		return []any{obj.pendingEndorsementCount}, nil
	default:
		return []any{}, fmt.Errorf("can't get field \"%s\"", path[0])
	}
}
