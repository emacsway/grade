package endorser

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type EndorserExporter struct {
	Id                        member.TenantMemberIdExporter
	Grade                     exporters.Uint8Exporter
	AvailableEndorsementCount exporters.UintExporter
	PendingEndorsementCount   exporters.UintExporter
	Version                   uint
	CreatedAt                 time.Time
}

func (ex *EndorserExporter) SetId(val member.TenantMemberId) {
	val.Export(&ex.Id)
}

func (ex *EndorserExporter) SetGrade(val grade.Grade) {
	val.Export(&ex.Grade)
}

func (ex *EndorserExporter) SetAvailableEndorsementCount(val EndorsementCount) {
	val.Export(&ex.AvailableEndorsementCount)
}

func (ex *EndorserExporter) SetPendingEndorsementCount(val EndorsementCount) {
	val.Export(&ex.PendingEndorsementCount)
}

func (ex *EndorserExporter) SetVersion(val uint) {
	ex.Version = val
}

func (ex *EndorserExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
