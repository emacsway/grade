package endorser

import (
	"database/sql/driver"
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/endorser"
	endorserVal "github.com/emacsway/grade/grade/internal/domain/endorser/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	s "github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	is "github.com/emacsway/grade/grade/internal/infrastructure/seedwork/specification"
)

type EndorserCanCompleteEndorsementSpecification struct {
	endorser.EndorserCanCompleteEndorsementSpecification
}

func (e *EndorserCanCompleteEndorsementSpecification) Compile() (sql string, params []driver.Valuer, err error) {
	exp := e.Expression()
	tv := is.NewTransformVisitor(GlobalScopeContext{})
	err = exp.Accept(tv)
	if err != nil {
		return "", nil, err
	}
	exp, err = tv.Result()
	if err != nil {
		return "", nil, err
	}
	v := is.NewPostgresqlVisitor()
	err = exp.Accept(v)
	if err != nil {
		return "", []driver.Valuer{}, err
	}
	return v.Result()
}

type EndorserContext struct {
}

func (c EndorserContext) AttrNode(parent s.EmptiableObject, path []string) (s.Visitable, error) {
	switch path[0] {
	case "availableEndorsementCount":
		return s.Field(parent, "available_endorsement_count"), nil
	case "pendingEndorsementCount":
		return s.Field(parent, "pending_endorsement_count"), nil
	case "id":
		return is.CompositeExpression(
			s.Field(parent, "tenant_id"),
			s.Field(parent, "member_id"),
		), nil
	default:
		return nil, fmt.Errorf("can't get field \"%s\"", path[0])
	}
}

type GlobalScopeContext struct {
	endorser EndorserContext
}

func (c GlobalScopeContext) AttrNode(path []string) (s.Visitable, error) {
	switch path[0] {
	case "endorser":
		return c.endorser.AttrNode(s.Object(s.GlobalScope(), "endorser"), path[1:])
	default:
		return nil, fmt.Errorf("can't get object \"%s\"", path[0])
	}
}

func (c GlobalScopeContext) ValueNode(val any) (s.Visitable, error) {
	switch valTyped := val.(type) {
	case endorserVal.EndorsementCount:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return s.Value(ex), nil
	case member.InternalMemberId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return s.Value(ex), nil
	case tenant.TenantId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return s.Value(ex), nil
	case member.MemberId:
		var ex MemberIdExporter
		valTyped.Export(&ex)
		nodes := []s.Visitable{}
		for _, v := range ex.Values() {
			node, err := c.ValueNode(v)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, node)
		}
		return is.CompositeExpression(nodes...), nil
	default:
		return nil, fmt.Errorf("can't export \"%#v\"", val)
	}
}

type MemberIdExporter struct {
	values [2]any
}

func (ex MemberIdExporter) Values() []any {
	return ex.values[:]
}

func (ex *MemberIdExporter) SetTenantId(val tenant.TenantId) {
	ex.values[0] = val
}

func (ex *MemberIdExporter) SetMemberId(val member.InternalMemberId) {
	ex.values[1] = val
}
