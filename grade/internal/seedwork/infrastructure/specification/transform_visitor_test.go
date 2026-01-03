package specification

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/seedwork/domain/identity"
	s "github.com/emacsway/grade/grade/internal/seedwork/domain/specification"
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
	sql string, params []any, err error,
) {
	return Compile(TestGlobalScopeContext{}, ss.Expression())
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

// Contexts

type SomethingScopeContext struct {
}

func (c SomethingScopeContext) AttrNode(parent s.EmptiableObject, path []string) (s.Visitable, error) {
	switch path[0] {
	case "id":
		return CompositeExpression(
			CompositeExpression(
				s.Field(parent, "tenant_id"),
				s.Field(parent, "member_id"),
			),
			s.Field(parent, "something_id"),
		), nil
	default:
		return nil, fmt.Errorf("can't get field \"%s\"", path[0])
	}
}

type TestGlobalScopeContext struct {
	something SomethingScopeContext
}

func (c TestGlobalScopeContext) AttrNode(path []string) (s.Visitable, error) {
	switch path[0] {
	case "something":
		return c.something.AttrNode(s.Object(s.GlobalScope(), "something"), path[1:])
	default:
		return nil, fmt.Errorf("can't get object \"%s\"", path[0])
	}
}

// FIXME: In case of stack implementation it will not work with member_id because this attrite is present on both cases:
// before transformation and after transformation.
// Нам нужно добавить JOIN в Visitor для Collection и создать alias для Item.
// Context нам нужно подменить, чтобы подменить наименование таблицы на alias of the JOIN.
// А в принципе, если весь Collection.expression запихнуть в JOIN ... ON, тогда alias может и не понадобится.
// Кажется, решение в том, чтобы выделить TransformContext с правилами преобразования.
// Нужно подумать что делать с полями сущностей 3-го и более глубокого уровня вложенности.
// В принципе, там должны получаться многоуровневые JOINs.
// Метод Extract() можно устранить, если значение возвращать тоже в контексте.
// Кажется, TransformVisitor можно вообще выбросить, т.к. сам контекст может возвращать CompositeExpression.
// Он все-равно управляет маппингом через err. Он создан для маппинга.

func (c TestGlobalScopeContext) ValueNode(val any) (s.Visitable, error) {
	switch valTyped := val.(type) {
	case InternalMemberId:
		var ex uint
		valTyped.Export(func(v uint) { ex = v })
		return s.Value(ex), nil
	case TenantId:
		var ex uint
		valTyped.Export(func(v uint) { ex = v })
		return s.Value(ex), nil
	case SomethingId:
		var ex uint
		valTyped.Export(func(v uint) { ex = v })
		return s.Value(ex), nil
	case MemberId:
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
		return CompositeExpression(nodes...), nil
	case MemberSomethingId:
		var ex MemberSomethingIdExporter
		valTyped.Export(&ex)
		nodes := []s.Visitable{}
		for _, v := range ex.Values() {
			node, err := c.ValueNode(v)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, node)
		}
		return CompositeExpression(nodes...), nil
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
	assert.Equal(t, 3, len(params))
	var tIdValue, mIdValue, sIdValue uint
	tId.Export(func(v uint) { tIdValue = v })
	mId.Export(func(v uint) { mIdValue = v })
	sId.Export(func(v uint) { sIdValue = v })

	assert.Equal(t, tIdValue, params[0])
	assert.Equal(t, mIdValue, params[1])
	assert.Equal(t, sIdValue, params[2])
}
