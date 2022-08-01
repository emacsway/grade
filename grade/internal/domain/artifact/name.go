package artifact

func NewName(val string) (Name, error) {
	return Name(val), nil
}

type Name string
