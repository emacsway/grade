package endorser

import (
	"fmt"

	s "github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
)

type EndorserCriteria struct{}

func (c EndorserCriteria) availableEndorsementCount() s.FieldNode {
	return s.Field(c.obj(), "availableEndorsementCount")
}

func (c EndorserCriteria) pendingEndorsementCount() s.FieldNode {
	return s.Field(c.obj(), "pendingEndorsementCount")
}

func (c EndorserCriteria) obj() s.ObjectNode {
	return s.Object(s.EmptyObject(), "endorser")
}

var endorser = EndorserCriteria{}

type EndorserCanCompleteEndorsementSpecification struct {
}

func (e EndorserCanCompleteEndorsementSpecification) Expression() s.Visitable {
	return s.And(
		s.NotEqual(
			endorser.availableEndorsementCount(),
			s.Value(EndorsementCount(0)),
		),
		s.NotEqual(
			endorser.pendingEndorsementCount(),
			s.Value(EndorsementCount(0)),
		),
		s.GreaterThanEqual(
			endorser.availableEndorsementCount(),
			endorser.pendingEndorsementCount(),
		),
	)
}

func (e EndorserCanCompleteEndorsementSpecification) IsSatisfiedBy(obj Endorser) (bool, error) {
	v := s.NewEvaluateVisitor(Context{
		endorser: obj,
	})
	err := e.Expression().Accept(v)
	if err != nil {
		return false, err
	}
	return v.Result()
}

type Context struct {
	endorser Endorser
}

func (c Context) ValuesByPath(path ...string) ([]any, error) {
	switch path[0] {
	case "endorser":
		return c.endorserPath(c.endorser, path[1:]...)
	default:
		return []any{}, fmt.Errorf("can't get object \"%s\"", path[0])
	}
}

func (c Context) endorserPath(obj Endorser, path ...string) ([]any, error) {
	switch path[0] {
	case "availableEndorsementCount":
		return []any{obj.availableEndorsementCount}, nil
	case "pendingEndorsementCount":
		return []any{obj.pendingEndorsementCount}, nil
	default:
		return []any{}, fmt.Errorf("can't get field \"%s\"", path[0])
	}
}
