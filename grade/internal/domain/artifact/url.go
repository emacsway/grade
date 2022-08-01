package artifact

func NewUrl(val string) (Url, error) {
	return Url(val), nil
}

type Url string
