package values

func NewFullNameExporter(firstName, lastName string) FullNameExporter {
	return FullNameExporter{
		FirstName: firstName,
		LastName:  lastName,
	}
}

type FullNameExporter struct {
	FirstName string
	LastName  string
}

func (ex *FullNameExporter) SetFirstName(val FirstName) {
	val.Export(func(v string) { ex.FirstName = v })
}

func (ex *FullNameExporter) SetLastName(val LastName) {
	val.Export(func(v string) { ex.LastName = v })
}
