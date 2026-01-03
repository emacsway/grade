package tenant

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

type TenantExporter struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	Version   uint
}

func (ex *TenantExporter) SetId(val values.TenantId) {
	val.Export(func(v uint) { ex.Id = v })
}

func (ex *TenantExporter) SetName(val values.Name) {
	val.Export(func(v string) { ex.Name = v })
}

func (ex *TenantExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}

func (ex *TenantExporter) SetVersion(val uint) {
	ex.Version = val
}
