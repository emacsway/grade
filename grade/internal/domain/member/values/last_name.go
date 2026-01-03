package values

func NewLastName(val string) (LastName, error) {
	return LastName(val), nil
}

type LastName string

func (n LastName) Export(ex func(string)) {
	ex(string(n))
}
