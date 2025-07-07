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
	values, err := v.Context.ValuesByPath(ExtractFieldPath(n)...)
	if err != nil {
		return err
	}
	v.SetCurrentValue(values)
	return nil
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
	lefts := v.CurrentValue()
	err = n.Right().Accept(v)
	if err != nil {
		return err
	}
	rights := v.CurrentValue()
	if v.yieldBooleanOperator(n.Operator()) {
		result := false
		// FIXME: здесь мы ищем совпадение по атрибуту любой из вложенных сущностей,
		// в то время как PostgresqlVisitor ищет совпадение по одной из вложенных сущностей.
		// В качестве решения можно воспользоваться относительным путем по аналогии @ в jsonpath.
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

func ExtractFieldPath(n FieldNode) []string {
	path := []string{n.Name()}
	var obj EmptiableObject = n.Object()
	for !obj.IsEmpty() {
		path = append([]string{obj.Name()}, path...)
		obj = obj.Parent()
	}
	return path
}
