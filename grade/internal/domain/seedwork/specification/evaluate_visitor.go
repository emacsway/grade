package specification

import (
	"errors"
	"fmt"
)

func NewEvaluateVisitor(context Context) *EvaluateVisitor {
	return &EvaluateVisitor{
		Context: context,
	}
}

type EvaluateVisitor struct {
	currentValue any
	currentItem  Context
	stack        []Context
	Context
}

func (v *EvaluateVisitor) push(ctx Context) {
	v.stack = append(v.stack, v.Context)
	v.Context = ctx
}

func (v *EvaluateVisitor) pop() {
	v.Context = v.stack[len(v.stack)-1]
	v.stack = v.stack[:len(v.stack)-1]
}

func (v EvaluateVisitor) CurrentValue() any {
	return v.currentValue
}

func (v *EvaluateVisitor) SetCurrentValue(val any) {
	v.currentValue = val
}

func (v *EvaluateVisitor) VisitGlobalScope(n GlobalScopeNode) error {
	v.push(v.Context)
	return nil
}

func (v *EvaluateVisitor) VisitObject(n ObjectNode) error {
	err := n.Parent().Accept(v)
	if err != nil {
		return err
	}
	obj, err := v.Context.Get(n.Name())
	v.pop()
	if err != nil {
		return err
	}
	v.push(obj.(Context))
	return nil
}

func (v *EvaluateVisitor) VisitCollection(n CollectionNode) error {
	err := n.Parent().Accept(v)
	if err != nil {
		return err
	}
	items, err := v.Context.Get(n.Name())
	v.pop()
	if err != nil {
		return err
	}
	itemsTyped, ok := items.([]Context)
	if !ok {
		return errors.New("currentValue is not a collection of Contexts")
	}
	result := false
	for i := range itemsTyped {
		v.currentItem = itemsTyped[i]
		err := n.Predicate().Accept(v)
		if err != nil {
			return err
		}
		result = result || v.CurrentValue().(bool)
	}
	v.SetCurrentValue(result)
	return nil
}

func (v *EvaluateVisitor) VisitItem(n ItemNode) error {
	v.push(v.currentItem)
	return nil
}

func (v *EvaluateVisitor) VisitField(n FieldNode) error {
	err := n.Object().Accept(v)
	if err != nil {
		return err
	}
	value, err := v.Context.Get(n.Name())
	v.pop()
	if err != nil {
		return err
	}
	v.SetCurrentValue(value)
	return nil
}

func (v *EvaluateVisitor) VisitValue(n ValueNode) error {
	v.SetCurrentValue(n.Value())
	return nil
}

func (v *EvaluateVisitor) VisitPrefix(n PrefixNode) error {
	err := n.Operand().Accept(v)
	if err != nil {
		return err
	}
	operand := v.CurrentValue()
	if v.yieldBooleanOperator(n.Operator()) {
		result, err := v.evalYieldBooleanPrefix(operand, n.Operator())
		if err != nil {
			return err
		}
		v.SetCurrentValue(result)
	} else {
		return fmt.Errorf("mathematical operator \"%s\" is not supported", n.Operator())
	}
	return nil
}
func (v EvaluateVisitor) evalYieldBooleanPrefix(operand any, op Operator) (bool, error) {
	switch op {
	case OperatorNot:
		return v.evalNot(operand)
	default:
		return false, fmt.Errorf("operator \"%s\" is not supported", op)
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
	left := v.CurrentValue()
	err = n.Right().Accept(v)
	if err != nil {
		return err
	}
	right := v.CurrentValue()
	if v.yieldBooleanOperator(n.Operator()) {
		result, err := v.evalYieldBooleanInfix(left, n.Operator(), right)
		if err != nil {
			return err
		}
		v.SetCurrentValue(result)
	} else {
		return fmt.Errorf("mathematical operator \"%s\" is not supported", n.Operator())
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
		return false, fmt.Errorf("operator \"%s\" is not supported", op)
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
	result := v.CurrentValue()
	resultTyped, ok := result.(bool)
	if !ok {
		return false, errors.New("the result is not a bool")
	}
	return resultTyped, nil
}

type Context interface {
	Get(string) (any, error)
}

func ExtractFieldPath(n FieldNode) []string {
	path := []string{n.Name()}
	var obj EmptiableObject = n.Object()
	for !obj.IsRoot() {
		path = append([]string{obj.Name()}, path...)
		obj = obj.Parent()
	}
	return path
}

type CollectionContext struct {
	items []Context
}

func (c CollectionContext) Get(slice string) (any, error) {
	if slice == "*" {
		return c.items, nil
	}
	return nil, fmt.Errorf("unsupported slice type \"%s\"", slice)
}
