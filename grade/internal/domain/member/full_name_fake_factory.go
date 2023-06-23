package member

func NewFullNameFakeFactory() FullNameFakeFactory {
	return FullNameFakeFactory{
		FirstName: "FirstName1",
		LastName:  "LastName1",
	}
}

type FullNameFakeFactory struct {
	FirstName string
	LastName  string
}

func (f FullNameFakeFactory) Create() (FullName, error) {
	return NewFullName(f.FirstName, f.LastName)
}
