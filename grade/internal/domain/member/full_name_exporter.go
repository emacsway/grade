package member

import "github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"

func NewFullNameExporter(firstName, lastName string) FullNameExporter {
	return FullNameExporter{
		FirstName: exporters.StringExporter(firstName),
		LastName:  exporters.StringExporter(lastName),
	}
}

type FullNameExporter struct {
	FirstName exporters.StringExporter
	LastName  exporters.StringExporter
}

func (ex *FullNameExporter) SetFirstName(val FirstName) {
	val.Export(&ex.FirstName)
}

func (ex *FullNameExporter) SetLastName(val LastName) {
	val.Export(&ex.LastName)
}
