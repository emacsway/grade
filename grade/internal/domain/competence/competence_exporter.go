package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type CompetenceExporter struct {
	Id        values.CompetenceIdExporter
	Name      exporters.StringExporter
	OwnerId   member.MemberIdExporter
	CreatedAt time.Time
	Version   uint
}

func (ex *CompetenceExporter) SetId(val values.CompetenceId) {
	val.Export(&ex.Id)
}

func (ex *CompetenceExporter) SetName(val values.Name) {
	val.Export(&ex.Name)
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
