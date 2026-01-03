package values

func NewStatus(val uint8) (Status, error) {
	return Status(val), nil
}

type Status uint8

func (s Status) Export(ex func(uint8)) {
	ex(uint8(s))
}

const (
	Inactive = Status(0)
	Active   = Status(1)
)
