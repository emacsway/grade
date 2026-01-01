package specification

import (
	"errors"
	"fmt"

	s "github.com/emacsway/grade/grade/internal/seedwork/domain/specification"
)

var (
	ErrCompositeExpressionsDifferentLength = errors.New("composite expressions have different length")
)

type Context interface {
	AttrNode(path []string) (s.Visitable, error)
	ValueNode(val any) (s.Visitable, error)
	// TODO: с вложенными контекстами ValueNode не будет работать, т.к. ValueNode может идти первым операндом. Нужно разделять интерфейсы.
}

func NewTransformVisitor(context Context) *TransformVisitor {
	return &TransformVisitor{
		Context: context,
	}
}

type TransformVisitor struct {
	currentNode s.Visitable
	stack       []Context
	Context
}

func (v *TransformVisitor) Push(ctx Context) {
	v.stack = append(v.stack, v.Context)
	v.Context = ctx
}

func (v *TransformVisitor) Pop() {
	v.Context = v.stack[len(v.stack)-1]
	v.stack = v.stack[:len(v.stack)-1]
}

func (v *TransformVisitor) VisitGlobalScope(_ s.GlobalScopeNode) error {
	// v.push(v.Context)
	return nil
}

func (v *TransformVisitor) VisitObject(_ s.ObjectNode) error {
	return nil
}

func (v *TransformVisitor) VisitCollection(n s.CollectionNode) error {
	return nil
}

func (v *TransformVisitor) VisitItem(n s.ItemNode) error {
	// v.push(v.currentItem)
	return nil
}

func (v *TransformVisitor) VisitField(n s.FieldNode) error {
	node, err := v.Context.AttrNode(s.ExtractFieldPath(n))
	// v.pop()
	if err != nil {
		return err
	}
	v.currentNode = node
	return nil
}

func (v *TransformVisitor) VisitValue(n s.ValueNode) error {
	node, err := v.Context.ValueNode(n.Value())
	if err != nil {
		return err
	}
	v.currentNode = node
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
	leftComposite, ok := left.(CompositeExpressionNode)
	if ok {
		rightComposite, ok := right.(CompositeExpressionNode)
		if !ok {
			return errors.New("not enough composite expressions")
		}
		switch n.Operator() {
		case s.OperatorEq:
			v.currentNode, err = leftComposite.Equal(rightComposite)
		case s.OperatorNe:
			v.currentNode, err = leftComposite.Equal(rightComposite)
		default:
			return fmt.Errorf("operator \"%s\" is not supported for composite expressions", n.Operator())
		}
		if err != nil {
			return err
		}
	} else {
		v.currentNode = s.NewInfixNode(left, n.Operator(), right, n.Associativity())
	}
	return nil
}

func (v TransformVisitor) Result() (s.Visitable, error) {
	return v.currentNode, nil
}
