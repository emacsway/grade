package tenant

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
)

type TenantExporter struct {
	Id        exporters.UintExporter
	Name      exporters.StringExporter
	CreatedAt time.Time
	Version   uint
}

func (ex *TenantExporter) SetId(val values.TenantId) {
	val.Export(&ex.Id)
}

func (ex *TenantExporter) SetName(val values.Name) {
	val.Export(&ex.Name)
}

func (ex *TenantExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}

func (ex *TenantExporter) SetVersion(val uint) {
	ex.Version = val
}
