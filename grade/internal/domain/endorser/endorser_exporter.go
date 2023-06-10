package recognizer

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type RecognizerExporter struct {
	Id                        member.TenantMemberIdExporter
	Grade                     exporters.Uint8Exporter
	AvailableEndorsementCount exporters.UintExporter
	PendingEndorsementCount   exporters.UintExporter
	Version                   uint
	CreatedAt                 time.Time
}

func (ex *RecognizerExporter) SetId(val member.TenantMemberId) {
	val.Export(&ex.Id)
}

func (ex *RecognizerExporter) SetGrade(val grade.Grade) {
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
