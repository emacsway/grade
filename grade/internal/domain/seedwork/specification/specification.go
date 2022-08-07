package specification

import (
	"errors"
	"fmt"
)

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

func Object(name string) ObjectNode {
	return ObjectNode{
		name: name,
	}
}

type ObjectNode struct {
	parent *ObjectNode
	name   string
}

func (n ObjectNode) Parent() *ObjectNode {
	return n.parent
}

func (n ObjectNode) Name() string {
	return n.name
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

func NewEvaluateVisitor(context Context) *EvaluateVisitor {
	return &EvaluateVisitor{
		Context: context,
	}
}

type EvaluateVisitor struct {
	currentValue []any
	Context
}

func (v EvaluateVisitor) CurrentValue() []any {
	return v.currentValue
}

func (v *EvaluateVisitor) SetCurrentValue(val []any) {
	v.currentValue = val
}

func (v *EvaluateVisitor) VisitObject(_ ObjectNode) error {
	// Is not used in Evaluation - only in SQL-building
	return nil
}

func (v *EvaluateVisitor) VisitField(n FieldNode) error {
	values, err := v.Context.ValuesByPath(v.extractFieldPath(n)...)
	if err != nil {
		return err
	}
	v.SetCurrentValue(values)
	return nil
}

func (v *EvaluateVisitor) extractFieldPath(n FieldNode) []string {
	path := []string{n.Name()}
	fistObj := n.Object()
	obj := &fistObj
	for obj != nil {
		path = append([]string{obj.Name()}, path...)
		obj = obj.Parent()
	}
	return path
}

func (v *EvaluateVisitor) VisitValue(n ValueNode) error {
	v.SetCurrentValue([]any{n.Value()})
	return nil
}

func (v *EvaluateVisitor) VisitPrefix(n PrefixNode) error {
	err := n.Operand().Accept(v)
	if err != nil {
		return err
	}
	operands := v.CurrentValue()
	if v.yieldBooleanOperator(n.Operator()) {
		// aggregate.[]entity.field bool
		result := false
		for i := range operands {
			nextResult, err := v.evalYieldBooleanPrefix(operands[i], n.Operator())
			if err != nil {
				return err
			}
			result = result || nextResult
		}
		v.SetCurrentValue([]any{result})
	} else {
		return fmt.Errorf("mathematical operator \"%s\" is not supperted", n.Operator())
	}
	return nil
}
func (v EvaluateVisitor) evalYieldBooleanPrefix(operand any, op Operator) (bool, error) {
	switch op {
	case OperatorNot:
		return v.evalNot(operand)
	default:
		return false, fmt.Errorf("operator \"%s\" is not supperted", op)
	}
}

func (v EvaluateVisitor) evalNot(operand any) (bool, error) {
	operandTyped, ok := operand.(bool)
	if !ok {
		return false, errors.New("operand is not a bool")
	}
	return !operandTyped, nil
}

func (v *EvaluateVisitor) VisitInfix(n InfixNode) error {
	err := n.Left().Accept(v)
	if err != nil {
		return err
	}
	lefts := v.CurrentValue()
	err = n.Right().Accept(v)
	if err != nil {
		return err
	}
	rights := v.CurrentValue()
	if v.yieldBooleanOperator(n.Operator()) {
		result := false
		for i := range lefts {
			for j := range rights {
				// aggregate.[]entity.field int == aggregate2.[]entity.field int
				nextResult, err := v.evalYieldBooleanInfix(lefts[i], n.Operator(), rights[j])
				if err != nil {
					return err
				}
				result = result || nextResult
			}
		}
		v.SetCurrentValue([]any{result})
	} else {
		return fmt.Errorf("mathematical operator \"%s\" is not supperted", n.Operator())
	}
	return nil
}

func (v EvaluateVisitor) yieldBooleanOperator(op Operator) bool {
	for i := range YieldBooleanOperators {
		if YieldBooleanOperators[i] == op {
			return true
		}
	}
	return false
}

func (v EvaluateVisitor) evalYieldBooleanInfix(left any, op Operator, right any) (bool, error) {
	switch op {
	case OperatorEq:
		return v.evalEq(left, right)
	case OperatorNe:
		return v.evalNe(left, right)
	case OperatorGt:
		return v.evalGt(left, right)
	case OperatorGte:
		return v.evalGte(left, right)
	case OperatorAnd:
		return v.evalAnd(left, right)
	default:
		return false, fmt.Errorf("operator \"%s\" is not supperted", op)
	}
}

func (v EvaluateVisitor) evalEq(left, right any) (bool, error) {
	leftTyped, ok := left.(EqualOperand)
	if !ok {
		return false, errors.New("left operand is not an EqualOperand")
	}
	rightTyped, ok := right.(EqualOperand)
	if !ok {
		return false, errors.New("right operand is not an EqualOperand")
	}
	return leftTyped.Equal(rightTyped), nil
}

func (v EvaluateVisitor) evalNe(left, right any) (bool, error) {
	leftTyped, ok := left.(EqualOperand)
	if !ok {
		return false, errors.New("left operand is not an EqualOperand")
	}
	rightTyped, ok := right.(EqualOperand)
	if !ok {
		return false, errors.New("right operand is not an EqualOperand")
	}
	return !leftTyped.Equal(rightTyped), nil
}

func (v EvaluateVisitor) evalGt(left, right any) (bool, error) {
	leftTyped, ok := left.(GreaterThanOperand)
	if !ok {
		return false, errors.New("left operand is not a GreaterThanOperand")
	}
	rightTyped, ok := right.(GreaterThanOperand)
	if !ok {
		return false, errors.New("right operand is not a GreaterThanOperand")
	}
	return leftTyped.GreaterThan(rightTyped), nil
}

func (v EvaluateVisitor) evalGte(left, right any) (bool, error) {
	leftTyped, ok := left.(GreaterThanEqualOperand)
	if !ok {
		return false, errors.New("left operand is not a GreaterThanEqualOperand")
	}
	rightTyped, ok := right.(GreaterThanEqualOperand)
	if !ok {
		return false, errors.New("right operand is not a GreaterThanEqualOperand")
	}
	return leftTyped.GreaterThanEqual(rightTyped), nil
}

func (v EvaluateVisitor) evalAnd(left, right any) (bool, error) {
	leftTyped, ok := left.(bool)
	if !ok {
		return false, errors.New("left operand is not a bool")
	}
	rightTyped, ok := right.(bool)
	if !ok {
		return false, errors.New("right operand is not a bool")
	}
	return leftTyped && rightTyped, nil
}

func (v EvaluateVisitor) Result() (bool, error) {
	results := v.CurrentValue()
	for i := range results {
		resultTyped, ok := results[i].(bool)
		if !ok {
			return false, errors.New("the result is not a bool")
		}
		if resultTyped {
			return resultTyped, nil
		}
	}
	return false, nil
}

type Context interface {
	ValuesByPath(...string) ([]any, error)
}
