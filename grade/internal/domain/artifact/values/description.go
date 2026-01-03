package values

func NewDescription(val string) (Description, error) {
	return Description(val), nil
}

type Description string

func (d Description) Export(ex func(string)) {
	ex(string(d))
}
