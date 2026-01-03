package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
)

type MemberExporter struct {
	Id        values.MemberIdExporter
	Status    uint8
	FullName  values.FullNameExporter
	CreatedAt time.Time
	Version   uint
}

func (ex *MemberExporter) SetId(val values.MemberId) {
	val.Export(&ex.Id)
}

func (ex *MemberExporter) SetStatus(val values.Status) {
	val.Export(func(v uint8) { ex.Status = v })
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
