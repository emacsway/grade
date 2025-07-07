package endorser

import (
	"database/sql/driver"
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/endorser"
	endorserVal "github.com/emacsway/grade/grade/internal/domain/endorser/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	s "github.com/emacsway/grade/grade/internal/infrastructure/seedwork/specification"
)

type EndorserCanCompleteEndorsementSpecification struct {
	endorser.EndorserCanCompleteEndorsementSpecification
}

func (e *EndorserCanCompleteEndorsementSpecification) Compile() (sql string, params []driver.Valuer, err error) {
	v := s.NewPostgresqlVisitor(GlobalScopeContext{})
	err = e.Expression().Accept(v)
	if err != nil {
		return "", []driver.Valuer{}, err
	}
	return v.Result()
}

type EndorserIdContext struct {
}

func (c EndorserIdContext) NameByPath(path ...string) (string, error) {
	if len(path) == 0 {
		return "", s.NewMissingFieldsError("tenantId", "memberId")
	}
	switch path[0] {
	case "tenantId":
		return "tenant_id", nil
	case "memberId":
		return "member_id", nil
	default:
		return "", fmt.Errorf("can't get field \"%s\"", path[0])
	}
}

func (c EndorserIdContext) Extract(val any) (driver.Valuer, error) {
	switch valTyped := val.(type) {
	case member.InternalMemberId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return nil, nil
	case tenant.TenantId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return nil, nil
	case member.MemberId:
		var ex MemberIdExporter
		valTyped.Export(&ex)
		return nil, s.NewMissingValuesError(ex.Values()...)
	default:
		return nil, fmt.Errorf("can't export \"%#v\"", val)
	}
}

type EndorserContext struct {
	id EndorserIdContext
}

func (c EndorserContext) NameByPath(path ...string) (string, error) {
	prefix := "endorser"
	var name string
	var err error
	switch path[0] {
	case "availableEndorsementCount":
		name = "available_endorsement_count"
	case "pendingEndorsementCount":
		name = "pending_endorsement_count"
	case "id":
		name, err = c.id.NameByPath(path[1:]...)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("can't get field \"%s\"", path[0])
	}
	return prefix + "." + name, nil
}

func (c EndorserContext) Extract(val any) (driver.Valuer, error) {
	switch valTyped := val.(type) {
	case endorserVal.EndorsementCount:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return ex, nil
	default:
		return nil, fmt.Errorf("can't export \"%#v\"", val)
	}
}

type GlobalScopeContext struct {
	endorser EndorserContext
}

func (c GlobalScopeContext) NameByPath(path ...string) (string, error) {
	switch path[0] {
	case "endorser":
		return c.endorser.NameByPath(path[1:]...)
	default:
		return "", fmt.Errorf("can't get object \"%s\"", path[0])
	}
}

func (c GlobalScopeContext) Extract(val any) (driver.Valuer, error) {
	switch valTyped := val.(type) {
	case endorserVal.EndorsementCount:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return ex, nil
	case member.InternalMemberId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return nil, nil
	case tenant.TenantId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return nil, nil
	case member.MemberId:
		var ex MemberIdExporter
		valTyped.Export(&ex)
		return nil, s.NewMissingValuesError(ex.Values()...)
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
