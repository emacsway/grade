package specification

type Associativity string

const (
	LeftAssociative  Associativity = "LEFT"
	RightAssociative Associativity = "RIGHT"
	NonAssociative   Associativity = "NON"
)

type Operator string

const (
	// Comparison

	OperatorEq  Operator = "="
	OperatorGt  Operator = ">"
	OperatorLt  Operator = "<"
	OperatorGte Operator = ">="
	OperatorLte Operator = "<="
	OperatorNe  Operator = "!="

	// Logical operators

	OperatorAnd Operator = "AND"
	OperatorOr  Operator = "OR"
	OperatorNot Operator = "NOT"

	// Mathematical

	/*
		OperatorAdd Operator = "+"
		OperatorSub Operator = "-"
		OperatorMul Operator = "*"
		OperatorDiv Operator = "/"
		OperatorMod Operator = "%"
	*/

	OperatorPos Operator = "+"
	OperatorNeg Operator = "-"
)

var YieldBooleanOperators = []Operator{
	OperatorEq,
	OperatorGt,
	OperatorLt,
	OperatorGte,
	OperatorLte,
	OperatorNe,
	OperatorAnd,
	OperatorOr,
	OperatorNot,
}

type Operable interface {
	Associativity() Associativity
	Operator() Operator
}

type Visitable interface {
	Accept(Visitor) error
}

type Visitor interface {
	VisitObject(ObjectNode) error
	VisitField(FieldNode) error
	VisitValue(ValueNode) error
	VisitPrefix(PrefixNode) error
	VisitInfix(InfixNode) error
}

func Value(value any) ValueNode {
	return ValueNode{
		value: value,
	}
}

type ValueNode struct {
	value any
}

func (n ValueNode) Value() any {
	return n.value
}

func (n ValueNode) Accept(v Visitor) error {
	return v.VisitValue(n)
}

func Not(operand Visitable) PrefixNode {
	return PrefixNode{
		operator:      OperatorNot,
		operand:       operand,
		associativity: RightAssociative,
	}
}

func NewPrefixNode(operator Operator, operand Visitable, associativity Associativity) PrefixNode {
	return PrefixNode{
		operator:      operator,
		operand:       operand,
		associativity: associativity,
	}
}

type PrefixNode struct {
	operator      Operator
	operand       Visitable
	associativity Associativity
}

func (n PrefixNode) Operand() Visitable {
	return n.operand
}
func (n PrefixNode) Operator() Operator {
	return n.operator
}
func (n PrefixNode) Associativity() Associativity {
	return n.associativity
}
func (n PrefixNode) Accept(v Visitor) error {
	return v.VisitPrefix(n)
}

func Equal(left, right Visitable) InfixNode {
	return InfixNode{
		left:          left,
		operator:      OperatorEq,
		right:         right,
		associativity: NonAssociative,
	}
}

func NotEqual(left, right Visitable) InfixNode {
	return InfixNode{
		left:          left,
		operator:      OperatorNe,
		right:         right,
		associativity: NonAssociative,
	}
}

func GreaterThan(left, right Visitable) InfixNode {
	return InfixNode{
		left:          left,
		operator:      OperatorGt,
		right:         right,
		associativity: NonAssociative,
	}
}

func GreaterThanEqual(left, right Visitable) InfixNode {
	return InfixNode{
		left:          left,
		operator:      OperatorGte,
		right:         right,
		associativity: NonAssociative,
	}
}

func And(left Visitable, rights ...Visitable) InfixNode {
	left, right := foldRights(And, left, rights...)
	return InfixNode{
		left:          left,
		operator:      OperatorAnd,
		right:         right,
		associativity: LeftAssociative,
	}
}

func foldRights(
	aCallable func(Visitable, ...Visitable) InfixNode,
	aLeft Visitable,
	aRights ...Visitable,
) (left, right Visitable) {
	for len(aRights) > 1 {
		aLeft = aCallable(aLeft, aRights[0])
		aRights = aRights[1:]
	}
	return aLeft, aRights[0]
}

func NewInfixNode(left Visitable, operator Operator, right Visitable, associativity Associativity) InfixNode {
	return InfixNode{
		left:          left,
		operator:      operator,
		right:         right,
		associativity: associativity,
	}
}

type InfixNode struct {
	left          Visitable
	operator      Operator
	right         Visitable
	associativity Associativity
}

func (n InfixNode) Left() Visitable {
	return n.left
}

func (n InfixNode) Operator() Operator {
	return n.operator
}

func (n InfixNode) Right() Visitable {
	return n.right
}

func (n InfixNode) Associativity() Associativity {
	return n.associativity
}

func (n InfixNode) Accept(v Visitor) error {
	return v.VisitInfix(n)
}

type EmptiableObject interface {
	Visitable
	Parent() EmptiableObject
	Name() string
	IsEmpty() bool
}

func EmptyObject() EmptyObjectNode {
	return EmptyObjectNode{}
}

type EmptyObjectNode struct{}

func (n EmptyObjectNode) Parent() EmptiableObject {
	return n
}

func (n EmptyObjectNode) Name() string {
	return "Empty"
}

func (n EmptyObjectNode) IsEmpty() bool {
	return true
}
func (n EmptyObjectNode) Accept(_ Visitor) error {
	return nil
}

func Object(parent EmptiableObject, name string) ObjectNode {
	return ObjectNode{
		parent: parent,
		name:   name,
	}
}

type ObjectNode struct {
	parent EmptiableObject
	name   string
}

func (n ObjectNode) Parent() EmptiableObject {
	return n.parent
}

func (n ObjectNode) Name() string {
	return n.name
}

func (n ObjectNode) IsEmpty() bool {
	return false
}

func (n ObjectNode) Accept(v Visitor) error {
	return v.VisitObject(n)
}

func Field(object ObjectNode, name string) FieldNode {
	return FieldNode{
		object: object,
		name:   name,
	}
}

type FieldNode struct {
	object ObjectNode
	name   string
}

func (n FieldNode) Name() string {
	return n.name
}

func (n FieldNode) Object() ObjectNode {
	return n.object
}

func (n FieldNode) Accept(v Visitor) error {
	return v.VisitField(n)
}
