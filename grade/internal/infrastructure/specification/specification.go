package specification

import (
	"database/sql/driver"
	"fmt"

	s "github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
)

type SqlVisitor struct {
	sql        string
	parameters []driver.Valuer
	Context
}

func (v *SqlVisitor) VisitObject(_ s.ObjectNode) error {
	return nil
}

func (v *SqlVisitor) VisitField(n s.FieldNode) error {
	name, err := v.Context.NameByPath(v.extractFieldPath(n)...)
	if err != nil {
		return err
	}
	v.sql += name
	return nil
}

func (v *SqlVisitor) extractFieldPath(n s.FieldNode) []string {
	path := []string{n.Name()}
	fistObj := n.Object()
	obj := &fistObj
	for obj != nil {
		path = append([]string{obj.Name()}, path...)
		obj = obj.Parent()
	}
	return path
}

func (v *SqlVisitor) VisitValue(n s.ValueNode) error {
	v.sql += "?"
	val, err := v.Extract(n.Value())
	if err != nil {
		return err
	}
	v.parameters = append(v.parameters, val)
	return nil
}

func (v *SqlVisitor) VisitInfix(n s.InfixNode) error {
	v.sql += "("
	err := n.Left().Accept(v)
	if err != nil {
		return err
	}
	v.sql += fmt.Sprintf(" %s ", n.Operator())
	err = n.Right().Accept(v)
	if err != nil {
		return err
	}
	v.sql += ")"
	return nil
}

func (v SqlVisitor) Result() (sql string, params []driver.Valuer, err error) {
	return v.sql, v.parameters, nil
}

type Context interface {
	NameByPath(...string) (string, error)
	Extract(any) (driver.Valuer, error)
}
