package faker

import (
	faker2 "syreclabs.com/go/faker"

	"github.com/icrowley/fake"
)

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

func (f Faker) Artifact() string {
	return fake.ProductName()
}

func (f Faker) Sentences() string {
	return fake.Sentences()
}

func (f Faker) Url() string {
	return faker2.Internet().Url()
}
