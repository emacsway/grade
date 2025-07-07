package endorser

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/endorser/values"
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
	return s.Object(s.GlobalScope(), "endorser")
}

var endorser = EndorserCriteria{}

type EndorserCanCompleteEndorsementSpecification struct {
}

func (e EndorserCanCompleteEndorsementSpecification) Expression() s.Visitable {
	return s.And(
		s.NotEqual(
			endorser.availableEndorsementCount(),
			s.Value(values.EndorsementCount(0)),
		),
		s.NotEqual(
			endorser.pendingEndorsementCount(),
			s.Value(values.EndorsementCount(0)),
		),
		s.GreaterThanEqual(
			endorser.availableEndorsementCount(),
			endorser.pendingEndorsementCount(),
		),
	)
}

func (e EndorserCanCompleteEndorsementSpecification) IsSatisfiedBy(obj Endorser) (bool, error) {
	v := s.NewEvaluateVisitor(GlobalScopeContext{
		endorser: EndorserContext{
			obj: obj,
		},
	})
	err := e.Expression().Accept(v)
	if err != nil {
		return false, err
	}
	return v.Result()
}

type EndorserContext struct {
	obj Endorser
}

func (c EndorserContext) Get(path ...string) ([]any, error) {
	switch path[0] {
	case "availableEndorsementCount":
		return []any{c.obj.availableEndorsementCount}, nil
	case "pendingEndorsementCount":
		return []any{c.obj.pendingEndorsementCount}, nil
	default:
		return []any{}, fmt.Errorf("can't get field \"%s\"", path[0])
	}
}

// TODO: Rename me to GlobalScopeContext
type GlobalScopeContext struct {
	endorser EndorserContext
}

func (c GlobalScopeContext) Get(path ...string) ([]any, error) {
	switch path[0] {
	case "endorser":
		return []any{c.endorser}, nil
	default:
		return []any{}, fmt.Errorf("can't get object \"%s\"", path[0])
	}
}
