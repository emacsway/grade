package competence

import (
	"time"

	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type CompetenceExporter struct {
	Id        TenantCompetenceIdExporter
	Name      exporters.StringExporter
	OwnerId   member.TenantMemberIdExporter
	Version   uint
	CreatedAt time.Time
}

func (ex *CompetenceExporter) SetId(val TenantCompetenceId) {
	val.Export(&ex.Id)
}

func (ex *CompetenceExporter) SetName(val Name) {
	val.Export(&ex.Name)
}

func (ex *CompetenceExporter) SetOwnerId(val member.TenantMemberId) {
	val.Export(&ex.OwnerId)
}

func (ex *CompetenceExporter) SetVersion(val uint) {
	ex.Version = val
}

func (ex *CompetenceExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
