package specification

type EqualOperand interface {
	Equal(EqualOperand) bool
}

type LessThanOperand interface {
	LessThan(LessThanOperand) bool
}

type GreaterThanOperand interface {
	GreaterThan(GreaterThanOperand) bool
}

type GreaterThanEqualOperand interface {
	GreaterThanEqual(GreaterThanEqualOperand) bool
}

type LessThanEqualOperand interface {
	LessThanEqual(LessThanEqualOperand) bool
}
