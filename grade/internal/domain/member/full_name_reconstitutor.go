package member

type FullNameReconstitutor struct {
	FirstName string
	LastName  string
}

func (r FullNameReconstitutor) Reconstitute() (FullName, error) {
	return NewFullName(r.FirstName, r.LastName)
}
