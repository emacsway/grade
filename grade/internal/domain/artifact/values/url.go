package values

func NewUrl(val string) (Url, error) {
	return Url(val), nil
}

type Url string

func (u Url) Export(ex func(string)) {
	ex(string(u))
}
