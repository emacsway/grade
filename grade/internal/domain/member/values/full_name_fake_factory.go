package values

import "github.com/emacsway/grade/grade/internal/domain/seedwork/faker"

func NewFullNameFaker() FullNameFaker {
	aFaker := faker.NewFaker()
	return FullNameFaker{
		FirstName: aFaker.FirstName(),
		LastName:  aFaker.LastName(),
	}
}

type FullNameFaker struct {
	FirstName string
	LastName  string
}

func (f FullNameFaker) Create() (FullName, error) {
	return NewFullName(f.FirstName, f.LastName)
}
