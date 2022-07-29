package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"time"
)

type RecognizerExporter struct {
	Id                        member.TenantMemberIdExporterSetter
	Grade                     seedwork.ExporterSetter[uint8]
	AvailableEndorsementCount seedwork.ExporterSetter[uint8]
	PendingEndorsementCount   seedwork.ExporterSetter[uint8]
	Version                   uint
	CreatedAt                 time.Time
}

func (ex *RecognizerExporter) SetState(
	id member.TenantMemberIdExporterSetter,
	grade seedwork.ExporterSetter[uint8],
	availableEndorsementCount seedwork.ExporterSetter[uint8],
	pendingEndorsementCount seedwork.ExporterSetter[uint8],
	version uint,
	createdAt time.Time,
) {
	ex.Id = id
	ex.Grade = grade
	ex.AvailableEndorsementCount = availableEndorsementCount
	ex.PendingEndorsementCount = pendingEndorsementCount
	ex.Version = version
	ex.CreatedAt = createdAt
}

type RecognizerState struct {
	Id                        member.TenantMemberIdState
	Grade                     uint8
	AvailableEndorsementCount uint8
	PendingEndorsementCount   uint8
	Version                   uint
	CreatedAt                 time.Time
}
