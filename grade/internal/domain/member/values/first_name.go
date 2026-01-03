package values

func NewFirstName(val string) (FirstName, error) {
	return FirstName(val), nil
}

type FirstName string

func (n FirstName) Export(ex func(string)) {
	ex(string(n))
}
