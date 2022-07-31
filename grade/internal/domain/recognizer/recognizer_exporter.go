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
	AvailableEndorsementCount seedwork.UintExporter
	PendingEndorsementCount   seedwork.UintExporter
	Version                   uint
	CreatedAt                 time.Time
}

func (ex *RecognizerExporter) SetId(val member.TenantMemberId) {
	val.Export(&ex.Id)
}

func (ex *RecognizerExporter) SetGrade(val shared.Grade) {
	val.Export(&ex.Grade)
}

func (ex *RecognizerExporter) SetAvailableEndorsementCount(val EndorsementCount) {
	val.Export(&ex.AvailableEndorsementCount)
}

func (ex *RecognizerExporter) SetPendingEndorsementCount(val EndorsementCount) {
	val.Export(&ex.PendingEndorsementCount)
}

func (ex *RecognizerExporter) SetVersion(val uint) {
	ex.Version = val
}

func (ex *RecognizerExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
