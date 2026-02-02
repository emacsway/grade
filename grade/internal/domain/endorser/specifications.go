package endorser

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/endorser/values"
	s "github.com/krew-solutions/ascetic-ddd-go/asceticddd/specification/domain"
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

func (c EndorserContext) Get(attr string) (any, error) {
	switch attr {
	case "availableEndorsementCount":
		return c.obj.availableEndorsementCount, nil
	case "pendingEndorsementCount":
		return c.obj.pendingEndorsementCount, nil
	default:
		return nil, fmt.Errorf("can't get field \"%s\"", attr)
	}
}

type GlobalScopeContext struct {
	endorser EndorserContext
}

func (c GlobalScopeContext) Get(attr string) (any, error) {
	switch attr {
	case "endorser":
		return c.endorser, nil
	default:
		return nil, fmt.Errorf("can't get object \"%s\"", attr)
	}
}
