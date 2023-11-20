package endorser

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/endorser/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type EndorserExporter struct {
	Id                        member.MemberIdExporter
	Grade                     exporters.Uint8Exporter
	AvailableEndorsementCount exporters.UintExporter
	PendingEndorsementCount   exporters.UintExporter
	CreatedAt                 time.Time
	Version                   uint
}

func (ex *EndorserExporter) SetId(val member.MemberId) {
	val.Export(&ex.Id)
}

func (ex *EndorserExporter) SetGrade(val grade.Grade) {
	val.Export(&ex.Grade)
}

func (ex *EndorserExporter) SetAvailableEndorsementCount(val values.EndorsementCount) {
	val.Export(&ex.AvailableEndorsementCount)
}

func (ex *EndorserExporter) SetPendingEndorsementCount(val values.EndorsementCount) {
	val.Export(&ex.PendingEndorsementCount)
}

func (ex *EndorserExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}

func (ex *EndorserExporter) SetVersion(val uint) {
	ex.Version = val
}
