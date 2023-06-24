package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type MemberExporter struct {
	Id        TenantMemberIdExporter
	Status    exporters.Uint8Exporter
	FullName  FullNameExporter
	Version   uint
	CreatedAt time.Time
}

func (ex *MemberExporter) SetId(val TenantMemberId) {
	val.Export(&ex.Id)
}

func (ex *MemberExporter) SetStatus(val Status) {
	val.Export(&ex.Status)
}

func (ex *MemberExporter) SetFullName(val FullName) {
	val.Export(&ex.FullName)
}

func (ex *MemberExporter) SetVersion(val uint) {
	ex.Version = val
}

func (ex *MemberExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
