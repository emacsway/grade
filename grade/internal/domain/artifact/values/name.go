package values

func NewName(val string) (Name, error) {
	return Name(val), nil
}

type Name string

func (n Name) Export(ex func(string)) {
	ex(string(n))
}
