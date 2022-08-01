package artifact

func NewDescription(val string) (Description, error) {
	return Description(val), nil
}

type Description string
