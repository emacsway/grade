package specification

import (
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
	s "github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

type SomethingCriteria struct {
}

func (sc SomethingCriteria) Id() s.FieldNode {
	return s.Field(sc.obj(), "id")
}

func (sc SomethingCriteria) obj() s.ObjectNode {
	return s.Object(s.EmptyObject(), "something")
}

type SomethingSpecification struct {
}

var something = SomethingCriteria{}
var tId, _ = identity.NewUuidIdentity(uuid.NewUuid())
var mId, _ = identity.NewUuidIdentity(uuid.NewUuid())
var sId, _ = identity.NewUuidIdentity(uuid.NewUuid())

func (ss SomethingSpecification) Expression() s.Visitable {
	return s.Equal(
		something.Id(),
		s.Value(MemberSomethingId{
			TenantMemberId{
				TenantId{tId},
				MemberId{mId},
			},
			SomethingId{
				sId,
			},
		}),
	)
}

func (ss SomethingSpecification) Execute( /* session infrastructure.PgxSession */ ) (
	sql string, params []driver.Valuer, err error,
) {
	exp := ss.Expression()
	for i := 1; i <= 10; i++ {
		v := NewTransformVisitor(TestContext{})
		err := exp.Accept(v)
		if err != nil {
			return "", nil, err
		}
		exp, err = v.Result()
		if err != nil {
			return "", nil, err
		}
		if !v.IsChanged() {
			break
		}
	}
	v := NewPostgresqlVisitor(TestContext{})
	err = exp.Accept(v)
	if err != nil {
		return "", nil, err
	}
	return v.Result()
}

type MemberSomethingId struct {
	memberId    TenantMemberId
	somethingId SomethingId
}

func (cid MemberSomethingId) Export(ex MemberSomethingIdExporterSetter) {
	ex.SetMemberId(cid.memberId)
	ex.SetSomethingId(cid.somethingId)
}

type MemberSomethingIdExporterSetter interface {
	SetMemberId(TenantMemberId)
	SetSomethingId(SomethingId)
}

type TenantMemberId struct {
	tenantId TenantId
	memberId MemberId
}

func (cid TenantMemberId) Export(ex TenantMemberIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetMemberId(cid.memberId)
}

type TenantMemberIdExporterSetter interface {
	SetTenantId(TenantId)
	SetMemberId(MemberId)
}

type TenantId struct {
	identity.UuidIdentity
}

type MemberId struct {
	identity.UuidIdentity
}

type SomethingId struct {
	identity.UuidIdentity
}

type TestContext struct {
}

func (c TestContext) NameByPath(path ...string) (string, error) {
	switch path[0] {
	case "something":
		return c.somethingPath("something", path[1:]...)
	default:
		return "", fmt.Errorf("can't get object \"%s\"", path[0])
	}
}

func (c TestContext) somethingPath(prefix string, path ...string) (string, error) {
	switch path[0] {
	case "id":
		return c.somethingIdPath(prefix, path[1:]...)
	default:
		return "", fmt.Errorf("can't get field \"%s\"", path[0])
	}
}

func (c TestContext) somethingIdPath(prefix string, path ...string) (string, error) {
	if len(path) == 0 {
		return "", NewMissingFieldsError("memberId", "somethingId")
	}
	switch path[0] {
	case "memberId":
		return c.somethingIdMemberIdPath(prefix, path[1:]...)
	case "somethingId":
		return prefix + ".something_id", nil
	default:
		return "", fmt.Errorf("can't get field \"%s\"", path[0])
	}
}

func (c TestContext) somethingIdMemberIdPath(prefix string, path ...string) (string, error) {
	if len(path) == 0 {
		return "", NewMissingFieldsError("tenantId", "memberId")
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

func (c TestContext) Extract(val any) (driver.Valuer, error) {
	switch valTyped := val.(type) {
	case MemberId:
		var ex exporters.UuidExporter
		valTyped.Export(&ex)
		return ex, nil
	case TenantId:
		var ex exporters.UuidExporter
		valTyped.Export(&ex)
		return ex, nil
	case SomethingId:
		var ex exporters.UuidExporter
		valTyped.Export(&ex)
		return ex, nil
	case TenantMemberId:
		var ex TenantMemberIdExporter
		valTyped.Export(&ex)
		return nil, NewMissingValuesError(ex.Values()...)
	case MemberSomethingId:
		var ex MemberSomethingIdExporter
		valTyped.Export(&ex)
		return nil, NewMissingValuesError(ex.Values()...)
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

func (ex *TenantMemberIdExporter) SetTenantId(val TenantId) {
	ex.values[0] = val
}

func (ex *TenantMemberIdExporter) SetMemberId(val MemberId) {
	ex.values[1] = val
}

type MemberSomethingIdExporter struct {
	values [2]any
}

func (ex MemberSomethingIdExporter) Values() []any {
	return ex.values[:]
}

func (ex *MemberSomethingIdExporter) SetMemberId(val TenantMemberId) {
	ex.values[0] = val
}

func (ex *MemberSomethingIdExporter) SetSomethingId(val SomethingId) {
	ex.values[1] = val
}

func TestSomethingSpecification(t *testing.T) {
	ss := SomethingSpecification{}
	sql, params, err := ss.Execute()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(
		t,
		"something.tenant_id = ? AND something.member_id = ? AND something.something_id = ?",
		sql)
	assert.Equal(t, []driver.Valuer{
		exporters.UuidExporter(tId.Value()),
		exporters.UuidExporter(mId.Value()),
		exporters.UuidExporter(sId.Value()),
	}, params)
}
