package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type MemberExporter struct {
	Id        values.TenantMemberIdExporter
	Status    exporters.Uint8Exporter
	FullName  values.FullNameExporter
	Version   uint
	CreatedAt time.Time
}

func (ex *MemberExporter) SetId(val values.TenantMemberId) {
	val.Export(&ex.Id)
}

func (ex *MemberExporter) SetStatus(val values.Status) {
	val.Export(&ex.Status)
}

func (ex *MemberExporter) SetFullName(val values.FullName) {
	val.Export(&ex.FullName)
}

func (ex *MemberExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}

func (ex *MemberExporter) SetVersion(val uint) {
	ex.Version = val
}
