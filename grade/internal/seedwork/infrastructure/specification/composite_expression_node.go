package specification

import (
	s "github.com/emacsway/grade/grade/internal/seedwork/domain/specification"
)

type ExpressionComposer interface {
	Equal(other CompositeExpressionNode) (s.Visitable, error)
	NotEqual(other CompositeExpressionNode) (s.Visitable, error)
	s.Visitable
}

func CompositeExpression(nodes ...s.Visitable) CompositeExpressionNode {
	return CompositeExpressionNode{
		nodes: nodes,
	}
}

type CompositeExpressionNode struct {
	nodes []s.Visitable
}

func (n CompositeExpressionNode) Equal(other CompositeExpressionNode) (s.Visitable, error) {
	var operands []s.Visitable
	if len(n.nodes) != len(other.nodes) {
		return nil, ErrCompositeExpressionsDifferentLength
	}
	for i := range n.nodes {
		left, right := n.nodes[i], other.nodes[i]
		leftComposite, ok := left.(CompositeExpressionNode)
		if ok {
			rightComposite, ok := right.(CompositeExpressionNode)
			if !ok {
				return nil, ErrCompositeExpressionsDifferentLength
			}
			newNode, err := leftComposite.Equal(rightComposite)
			if err != nil {
				return nil, err
			}
			operands = append(operands, newNode)
		} else {
			operands = append(operands, s.Equal(left, right))
		}
	}
	return s.And(operands[0], operands[1:]...), nil
}

func (n CompositeExpressionNode) NotEqual(other CompositeExpressionNode) (s.Visitable, error) {
	var operands []s.Visitable
	if len(n.nodes) != len(other.nodes) {
		return nil, ErrCompositeExpressionsDifferentLength
	}
	for i := range n.nodes {
		left, right := n.nodes[i], other.nodes[i]
		leftComposite, ok := left.(CompositeExpressionNode)
		if ok {
			rightComposite, ok := right.(CompositeExpressionNode)
			if !ok {
				return nil, ErrCompositeExpressionsDifferentLength
			}
			newNode, err := leftComposite.NotEqual(rightComposite)
			if err != nil {
				return nil, err
			}
			operands = append(operands, newNode)
		} else {
			operands = append(operands, s.NotEqual(left, right))
		}
	}
	return s.Not(s.And(operands[0], operands[1:]...)), nil
}

func (n CompositeExpressionNode) Accept(v s.Visitor) error {
	return nil
}
