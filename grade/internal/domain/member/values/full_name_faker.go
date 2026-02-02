package values

import "github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/faker"

func NewFullNameFaker() FullNameFaker {
	f := FullNameFaker{}
	f.fake()
	return f
}

type FullNameFaker struct {
	FirstName string
	LastName  string
}

func (f *FullNameFaker) fake() {
	aFaker := faker.NewFaker()
	f.FirstName = aFaker.FirstName()
	f.LastName = aFaker.LastName()
}

func (f *FullNameFaker) Next() {
	f.fake()
}

func (f FullNameFaker) Create() (FullName, error) {
	return NewFullName(f.FirstName, f.LastName)
}
