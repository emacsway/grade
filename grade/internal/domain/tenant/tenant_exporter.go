package tenant

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type TenantExporter struct {
	Id        exporters.UintExporter
	Name      exporters.StringExporter
	Version   uint
	CreatedAt time.Time
}

func (ex *TenantExporter) SetId(val TenantId) {
	val.Export(&ex.Id)
}

func (ex *TenantExporter) SetName(val Name) {
	val.Export(&ex.Name)
}

func (ex *TenantExporter) SetVersion(val uint) {
	ex.Version = val
}

func (ex *TenantExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
