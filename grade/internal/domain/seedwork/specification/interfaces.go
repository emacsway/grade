package specification

type EqualOperand interface {
	Equal(EqualOperand) bool
}

type GreaterThanOperand interface {
	GreaterThan(GreaterThanOperand) bool
}

type GreaterThanEqualOperand interface {
	GreaterThanEqual(GreaterThanEqualOperand) bool
}
