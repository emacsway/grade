package values

func NewReason(reason string) (Reason, error) {
	return Reason(reason), nil
}

type Reason string

func (r Reason) Export(ex func(string)) {
	ex(string(r))
}
