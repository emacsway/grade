package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

type CompetenceExporter struct {
	Id        values.CompetenceIdExporter
	Name      string
	OwnerId   member.MemberIdExporter
	CreatedAt time.Time
	Version   uint
}

func (ex *CompetenceExporter) SetId(val values.CompetenceId) {
	val.Export(&ex.Id)
}

func (ex *CompetenceExporter) SetName(val values.Name) {
	val.Export(func(v string) { ex.Name = v })
}

func (ex *CompetenceExporter) SetOwnerId(val member.MemberId) {
	val.Export(&ex.OwnerId)
}

func (ex *CompetenceExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}

func (ex *CompetenceExporter) SetVersion(val uint) {
	ex.Version = val
}
