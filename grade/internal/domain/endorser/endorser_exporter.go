package endorser

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/endorser/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

type EndorserExporter struct {
	Id                        member.MemberIdExporter
	Grade                     uint8
	AvailableEndorsementCount uint
	PendingEndorsementCount   uint
	CreatedAt                 time.Time
	Version                   uint
}

func (ex *EndorserExporter) SetId(val member.MemberId) {
	val.Export(&ex.Id)
}

func (ex *EndorserExporter) SetGrade(val grade.Grade) {
	val.Export(func(v uint8) { ex.Grade = v })
}

func (ex *EndorserExporter) SetAvailableEndorsementCount(val values.EndorsementCount) {
	val.Export(func(v uint) { ex.AvailableEndorsementCount = v })
}

func (ex *EndorserExporter) SetPendingEndorsementCount(val values.EndorsementCount) {
	val.Export(func(v uint) { ex.PendingEndorsementCount = v })
}

func (ex *EndorserExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}

func (ex *EndorserExporter) SetVersion(val uint) {
	ex.Version = val
}
