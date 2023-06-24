package values

func NewFullName(firstName, lastName string) (FullName, error) {
	firstNameValue, err := NewFirstName(firstName)
	if err != nil {
		return FullName{}, err
	}
	lastNameValue, err := NewLastName(lastName)
	if err != nil {
		return FullName{}, err
	}
	return FullName{
		firstName: firstNameValue,
		lastName:  lastNameValue,
	}, nil
}

type FullName struct {
	firstName FirstName
	lastName  LastName
}

func (fn FullName) FirstName() FirstName {
	return fn.firstName
}

func (fn FullName) LastName() LastName {
	return fn.lastName
}

func (fn FullName) Export(ex FullNameExporterSetter) {
	ex.SetFirstName(fn.firstName)
	ex.SetLastName(fn.lastName)
}

type FullNameExporterSetter interface {
	SetFirstName(FirstName)
	SetLastName(LastName)
}
