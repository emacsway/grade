package specification

import (
	"errors"
	"fmt"

	s "github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
)

var (
	var ErrCompositeExpressionsDifferentLength = errors.New("composite expressions have different length")
)

func NewTransformVisitor(context Context) *TransformVisitor {
	return &TransformVisitor{
		Context: context,
	}
}

type TransformVisitor struct {
	compositeExpressions []CompositeExpression
	currentNode          s.Visitable
	Context
}

func (v *TransformVisitor) VisitObject(_ s.ObjectNode) error {
	return nil
}

func (v *TransformVisitor) VisitField(n s.FieldNode) error {
	_, err := v.Context.NameByPath(s.ExtractFieldPath(n)...)
	if err != nil {
		if errTyped, ok := err.(MissingFieldsError); ok {
			names := errTyped.MissingFieldNames()
			o := s.Object(n.Object(), n.Name())
			compositeExpression := CompositeExpression{}
			for i := range names {
				compositeExpression.Add(s.Field(o, names[i]))
			}
			return nil
		} else {
			return err
		}
	}
	v.currentNode = n
	return nil
}

func (v *TransformVisitor) VisitValue(n s.ValueNode) error {
	_, err := v.Extract(n.Value())
	if err != nil {
		return err
	}
	return nil
}

func (v *TransformVisitor) VisitPrefix(n s.PrefixNode) error {
	err := n.Operand().Accept(v)
	if err != nil {
		return err
	}
	v.currentNode = s.NewPrefixNode(n.Operator(), v.currentNode, n.Associativity())
	return nil
}

func (v *TransformVisitor) VisitInfix(n s.InfixNode) error {
	err := n.Left().Accept(v)
	if err != nil {
		return err
	}
	left := v.currentNode
	err = n.Right().Accept(v)
	if err != nil {
		return err
	}
	right := v.currentNode
	newNode, err := v.buildNodeFromCompositeExpressions(n)
	if err != nil {
		return err
	}
	if newNode != nil {
		v.currentNode = s.NewInfixNode(left, n.Operator(), right, n.Associativity())
	}
	return nil
}

func (v *TransformVisitor) buildNodeFromCompositeExpressions(n s.InfixNode) (s.Visitable, error) {
	l := len(v.compositeExpressions)
	if l != 0 {
		left, right := v.compositeExpressions[l-2], v.compositeExpressions[l-1]
		v.compositeExpressions = v.compositeExpressions[:l-2]
		switch n.Operator() {
		case s.OperatorEq:
			return left.Equal(right)
		case s.OperatorNe:
			return left.NotEqual(right)
		default:
			return nil, fmt.Errorf("operator \"%s\" is not supported for composite expressions", n.Operator())
		}
	}
	return nil, nil
}

func (v TransformVisitor) Result() (s.Visitable, error) {
	return v.currentNode, nil
}

type CompositeExpression struct {
	nodes []s.Visitable
}

func (n *CompositeExpression) Add(nodes ...s.Visitable) {
	n.nodes = append(n.nodes, nodes...)
}

func (n CompositeExpression) Equal(other CompositeExpression) (s.Visitable, error) {
	operands := []s.Visitable{}
	if len(n.nodes) != len(other.nodes) {
		return nil, ErrCompositeExpressionsDifferentLength
	}
	for i := range n.nodes {
		operands = append(operands, s.Equal(n.nodes[i], other.nodes[i]))
	}
	return s.And(operands[0], operands[1:]...), nil
}

func (n CompositeExpression) NotEqual(other CompositeExpression) (s.Visitable, error) {
	operands := []s.Visitable{}
	if len(n.nodes) != len(other.nodes) {
		return nil, ErrCompositeExpressionsDifferentLength
	}
	for i := range n.nodes {
		operands = append(operands, s.Equal(n.nodes[i], other.nodes[i]))
	}
	return s.Not(s.And(operands[0], operands[1:]...)), nil
}

func NewMissingFieldsError(names ...string) MissingFieldsError {
	return MissingFieldsError{
		missingFieldNames: names,
	}
}

type MissingFieldsError struct {
	missingFieldNames []string
}

func (e MissingFieldsError) MissingFieldNames() []string {
	return e.missingFieldNames
}

func (e MissingFieldsError) Error() string {
	return fmt.Sprintf("Missing names: %#v", e.missingFieldNames)
}

func NewMissingValuesError(values ...any) MissingValuesError {
	return MissingValuesError{
		missingValues: values,
	}
}

type MissingValuesError struct {
	missingValues []any
}

func (e MissingValuesError) MissingValues() []any {
	return e.missingValues
}

func (e MissingValuesError) Error() string {
	return fmt.Sprintf("Missing values: %#v", e.missingValues)
}
