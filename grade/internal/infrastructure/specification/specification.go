package specification

import (
	"database/sql/driver"
	"fmt"

	s "github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
)

func NewPostgresqlVisitor(context Context) *PostgresqlVisitor {
	v := &PostgresqlVisitor{
		precedenceMapping: make(map[string]int),
		Context:           context,
	}
	// https://www.postgresql.org/docs/14/sql-syntax-lexical.html#SQL-PRECEDENCE-TABLE
	v.setPrecedence(160, ". LEFT")
	v.setPrecedence(160, ":: LEFT")
	v.setPrecedence(150, "[ LEFT")
	v.setPrecedence(140, "+ RIGHT", "- RIGHT")
	v.setPrecedence(130, "^ LEFT")
	v.setPrecedence(120, "* LEFT", "/ LEFT", "% LEFT")
	v.setPrecedence(110, "+ LEFT", "- LEFT")
	// all other native and user-defined operators 👇️
	v.setPrecedence(100, "(any other operator) LEFT")
	v.setPrecedence(90, "BETWEEN NON", "IN NON", "LIKE NON", "ILIKE NON", "SIMILAR NON")
	v.setPrecedence(80, "< NON", "> NON", "= NON", "<= NON", ">= NON", "<> NON")
	v.setPrecedence(70, "IS NON", "ISNULL NON", "NOTNULL NON")
	v.setPrecedence(60, "NOT RIGHT")
	v.setPrecedence(50, "AND LEFT")
	v.setPrecedence(40, "OR LEFT")
	return v
}

type PostgresqlVisitor struct {
	sql               string
	parameters        []driver.Valuer
	precedence        int
	precedenceMapping map[string]int
	Context
}

func (v PostgresqlVisitor) getNodePrecedenceKey(n s.Operable) string {
	operator := n.Operator()
	return fmt.Sprintf("%s %s", operator, n.Associativity())
}
func (v PostgresqlVisitor) setPrecedence(precedence int, operators ...string) {
	for _, op := range operators {
		v.precedenceMapping[op] = precedence
	}
}

func (v *PostgresqlVisitor) visit(precedenceKey string, callable func() error) error {
	outerPrecedence := v.precedence
	innerPrecedence, ok := v.precedenceMapping[precedenceKey]
	if !ok {
		innerPrecedence, ok = v.precedenceMapping["(any other operator) LEFT"]
		if !ok {
			innerPrecedence = outerPrecedence
		}
	}
	v.precedence = innerPrecedence
	if innerPrecedence < outerPrecedence {
		v.sql += "("
	}
	err := callable()
	if err != nil {
		return err
	}
	if innerPrecedence < outerPrecedence {
		v.sql += ")"
	}
	v.precedence = outerPrecedence
	return nil
}

func (v *PostgresqlVisitor) VisitObject(_ s.ObjectNode) error {
	return nil
}

func (v *PostgresqlVisitor) VisitField(n s.FieldNode) error {
	name, err := v.Context.NameByPath(v.extractFieldPath(n)...)
	if err != nil {
		return err
	}
	v.sql += name
	return nil
}

func (v *PostgresqlVisitor) extractFieldPath(n s.FieldNode) []string {
	path := []string{n.Name()}
	fistObj := n.Object()
	obj := &fistObj
	for obj != nil {
		path = append([]string{obj.Name()}, path...)
		obj = obj.Parent()
	}
	return path
}

func (v *PostgresqlVisitor) VisitValue(n s.ValueNode) error {
	v.sql += "?"
	val, err := v.Extract(n.Value())
	if err != nil {
		return err
	}
	v.parameters = append(v.parameters, val)
	return nil
}

func (v *PostgresqlVisitor) VisitInfix(n s.InfixNode) error {
	precedenceKey := v.getNodePrecedenceKey(n)
	return v.visit(precedenceKey, func() error {
		err := n.Left().Accept(v)
		if err != nil {
			return err
		}
		v.sql += fmt.Sprintf(" %s ", n.Operator())
		err = n.Right().Accept(v)
		if err != nil {
			return err
		}
		return nil
	})
}

func (v PostgresqlVisitor) Result() (sql string, params []driver.Valuer, err error) {
	return v.sql, v.parameters, nil
}

type Context interface {
	NameByPath(...string) (string, error)
	Extract(any) (driver.Valuer, error)
}
