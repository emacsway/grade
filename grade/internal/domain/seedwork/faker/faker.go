package faker

import "github.com/icrowley/fake"

func NewFaker() Faker {
	return Faker{}
}

type Faker struct {
}

func (f Faker) Company() string {
	return fake.Company()
}

func (f Faker) FirstName() string {
	return fake.FirstName()
}

func (f Faker) LastName() string {
	return fake.LastName()
}

func (f Faker) Competence() string {
	return fake.Industry()
}
