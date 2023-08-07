package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competence "github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type ArtifactExporter struct {
	Id            values.TenantArtifactIdExporter
	Status        exporters.Uint8Exporter
	Name          exporters.StringExporter
	Description   exporters.StringExporter
	Url           exporters.StringExporter
	CompetenceIds []competence.TenantCompetenceIdExporter
	AuthorIds     []member.TenantMemberIdExporter
	OwnerId       member.TenantMemberIdExporter
	CreatedAt     time.Time
	Version       uint
}

func (ex *ArtifactExporter) SetId(val values.TenantArtifactId) {
	val.Export(&ex.Id)
}

func (ex *ArtifactExporter) SetStatus(val values.Status) {
	val.Export(&ex.Status)
}

func (ex *ArtifactExporter) SetName(val values.Name) {
	val.Export(&ex.Name)
}

func (ex *ArtifactExporter) SetDescription(val values.Description) {
	val.Export(&ex.Description)
}

func (ex *ArtifactExporter) SetUrl(val values.Url) {
	val.Export(&ex.Url)
}

func (ex *ArtifactExporter) AddCompetenceId(val competence.TenantCompetenceId) {
	var competenceExporter competence.TenantCompetenceIdExporter
	val.Export(&competenceExporter)
	ex.CompetenceIds = append(ex.CompetenceIds, competenceExporter)
}

func (ex *ArtifactExporter) AddAuthorId(val member.TenantMemberId) {
	var authorExporter member.TenantMemberIdExporter
	val.Export(&authorExporter)
	ex.AuthorIds = append(ex.AuthorIds, authorExporter)
}

func (ex *ArtifactExporter) SetOwnerId(val member.TenantMemberId) {
	val.Export(&ex.OwnerId)
}

func (ex *ArtifactExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}

func (ex *ArtifactExporter) SetVersion(val uint) {
	ex.Version = val
}
