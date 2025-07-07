package specification

import (
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
	s "github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
)

type SomethingCriteria struct {
}

func (sc SomethingCriteria) Id() s.FieldNode {
	return s.Field(sc.obj(), "id")
}

func (sc SomethingCriteria) obj() s.ObjectNode {
	return s.Object(s.GlobalScope(), "something")
}

type SomethingSpecification struct {
}

var something = SomethingCriteria{}
var tId, _ = identity.NewIntIdentity(10)
var mId, _ = identity.NewIntIdentity(3)
var sId, _ = identity.NewIntIdentity(3)

func (ss SomethingSpecification) Expression() s.Visitable {
	return s.Equal(
		something.Id(),
		s.Value(MemberSomethingId{
			MemberId{
				TenantId{tId},
				InternalMemberId{mId},
			},
			SomethingId{
				sId,
			},
		}),
	)
}

func (ss SomethingSpecification) Evaluate( /* session session.PgxSession */ ) (
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
	memberId    MemberId
	somethingId SomethingId
}

func (cid MemberSomethingId) Export(ex MemberSomethingIdExporterSetter) {
	ex.SetMemberId(cid.memberId)
	ex.SetSomethingId(cid.somethingId)
}

type MemberSomethingIdExporterSetter interface {
	SetMemberId(MemberId)
	SetSomethingId(SomethingId)
}

type MemberId struct {
	tenantId TenantId
	memberId InternalMemberId
}

func (cid MemberId) Export(ex MemberIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetMemberId(cid.memberId)
}

type MemberIdExporterSetter interface {
	SetTenantId(TenantId)
	SetMemberId(InternalMemberId)
}

type TenantId struct {
	identity.IntIdentity
}

type InternalMemberId struct {
	identity.IntIdentity
}

type SomethingId struct {
	identity.IntIdentity
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
	case InternalMemberId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return ex, nil
	case TenantId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return ex, nil
	case SomethingId:
		var ex exporters.UintExporter
		valTyped.Export(&ex)
		return ex, nil
	case MemberId:
		var ex MemberIdExporter
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

type MemberIdExporter struct {
	values [2]any
}

func (ex MemberIdExporter) Values() []any {
	return ex.values[:]
}

func (ex *MemberIdExporter) SetTenantId(val TenantId) {
	ex.values[0] = val
}

func (ex *MemberIdExporter) SetMemberId(val InternalMemberId) {
	ex.values[1] = val
}

type MemberSomethingIdExporter struct {
	values [2]any
}

func (ex MemberSomethingIdExporter) Values() []any {
	return ex.values[:]
}

func (ex *MemberSomethingIdExporter) SetMemberId(val MemberId) {
	ex.values[0] = val
}

func (ex *MemberSomethingIdExporter) SetSomethingId(val SomethingId) {
	ex.values[1] = val
}

func TestSomethingSpecification(t *testing.T) {
	ss := SomethingSpecification{}
	sql, params, err := ss.Evaluate()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(
		t,
		"something.tenant_id = $1 AND something.member_id = $2 AND something.something_id = $3",
		sql)
	assert.Equal(t, []driver.Valuer{
		exporters.UintExporter(tId.Value()),
		exporters.UintExporter(mId.Value()),
		exporters.UintExporter(sId.Value()),
	}, params)
}
