package recognizer

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

type RecognizerExporter struct {
	Id                        member.TenantMemberIdExporter
	Grade                     seedwork.Uint8Exporter
	AvailableEndorsementCount seedwork.ExporterSetter[uint8]
	PendingEndorsementCount   seedwork.ExporterSetter[uint8]
	Version                   uint
	CreatedAt                 time.Time
}

func (ex *RecognizerExporter) SetState(
	availableEndorsementCount seedwork.ExporterSetter[uint8],
	pendingEndorsementCount seedwork.ExporterSetter[uint8],
	version uint,
	createdAt time.Time,
) {
	ex.AvailableEndorsementCount = availableEndorsementCount
	ex.PendingEndorsementCount = pendingEndorsementCount
	ex.Version = version
	ex.CreatedAt = createdAt
}

func (ex *RecognizerExporter) SetId(id member.TenantMemberId) {
	id.Export(&ex.Id)
}

func (ex *RecognizerExporter) SetGrade(g shared.Grade) {
	g.Export(&ex.Grade)
}
