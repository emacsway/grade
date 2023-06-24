package member

import "github.com/emacsway/grade/grade/internal/domain/seedwork/faker"

func NewFullNameFakeFactory() FullNameFakeFactory {
	aFaker := faker.NewFaker()
	return FullNameFakeFactory{
		FirstName: aFaker.FirstName(),
		LastName:  aFaker.LastName(),
	}
}

type FullNameFakeFactory struct {
	FirstName string
	LastName  string
}

func (f FullNameFakeFactory) Create() (FullName, error) {
	return NewFullName(f.FirstName, f.LastName)
}
