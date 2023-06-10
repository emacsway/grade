package recognizer

import (
	"database/sql/driver"
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/recognizer"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
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
	case "id":
		return c.recognizerIdPath(prefix, path[1:]...)
	default:
		return "", fmt.Errorf("can't get field \"%s\"", path[0])
	}
}

func (c Context) recognizerIdPath(prefix string, path ...string) (string, error) {
	if len(path) == 0 {
		return "", s.NewMissingFieldsError("tenantId", "memberId")
	}
	switch path[0] {
	case "tenantId":
		return prefix + ".tenant_id", nil
	case "memberId":
		return prefix + ".member_id", nil
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
	case member.MemberId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return nil, nil
	case tenant.TenantId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return nil, nil
	case member.TenantMemberId:
		var ex TenantMemberIdExporter
		valTyped.Export(&ex)
		return nil, s.NewMissingValuesError(ex.Values()...)
	default:
		return nil, fmt.Errorf("can't export \"%#v\"", val)
	}
}

type TenantMemberIdExporter struct {
	values [2]any
}

func (ex TenantMemberIdExporter) Values() []any {
	return ex.values[:]
}

func (ex *TenantMemberIdExporter) SetTenantId(val tenant.TenantId) {
	ex.values[0] = val
}

func (ex *TenantMemberIdExporter) SetMemberId(val member.MemberId) {
	ex.values[1] = val
}
