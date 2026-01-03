package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competence "github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

type ArtifactExporter struct {
	Id            values.ArtifactIdExporter
	Status        uint8
	Name          string
	Description   string
	Url           string
	CompetenceIds []competence.CompetenceIdExporter
	AuthorIds     []member.MemberIdExporter
	OwnerId       member.MemberIdExporter
	CreatedAt     time.Time
	Version       uint
}

func (ex *ArtifactExporter) SetId(val values.ArtifactId) {
	val.Export(&ex.Id)
}

func (ex *ArtifactExporter) SetStatus(val values.Status) {
	val.Export(func(v uint8) { ex.Status = v })
}

func (ex *ArtifactExporter) SetName(val values.Name) {
	val.Export(func(v string) { ex.Name = v })
}

func (ex *ArtifactExporter) SetDescription(val values.Description) {
	val.Export(func(v string) { ex.Description = v })
}

func (ex *ArtifactExporter) SetUrl(val values.Url) {
	val.Export(func(v string) { ex.Url = v })
}

func (ex *ArtifactExporter) AddCompetenceId(val competence.CompetenceId) {
	var competenceExporter competence.CompetenceIdExporter
	val.Export(&competenceExporter)
	ex.CompetenceIds = append(ex.CompetenceIds, competenceExporter)
}

func (ex *ArtifactExporter) AddAuthorId(val member.MemberId) {
	var authorExporter member.MemberIdExporter
	val.Export(&authorExporter)
	ex.AuthorIds = append(ex.AuthorIds, authorExporter)
}

func (ex *ArtifactExporter) SetOwnerId(val member.MemberId) {
	val.Export(&ex.OwnerId)
}

func (ex *ArtifactExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}

func (ex *ArtifactExporter) SetVersion(val uint) {
	ex.Version = val
}
