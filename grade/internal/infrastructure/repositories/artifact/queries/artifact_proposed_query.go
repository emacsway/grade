package queries

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/seedwork/repository"
)

type ArtifactProposedQuery struct {
	repository.EventInsertQuery
	payload ArtifactProposedPayload
}

func (q *ArtifactProposedQuery) SetId(val values.TenantArtifactId) {
	val.Export(&q.payload.Id)
	val.Export(q)
}

func (q *ArtifactProposedQuery) SetArtifactId(val values.ArtifactId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.SetStreamId(v.String())
}

func (q *ArtifactProposedQuery) SetStatus(val values.Status) {
	val.Export(&q.payload.Status)
}
func (q *ArtifactProposedQuery) SetName(val values.Name) {
	val.Export(&q.payload.Name)
}
func (q *ArtifactProposedQuery) SetDescription(val values.Description) {
	val.Export(&q.payload.Description)
}
func (q *ArtifactProposedQuery) SetUrl(val values.Url) {
	val.Export(&q.payload.Url)
}
func (q *ArtifactProposedQuery) AddCompetenceId(val competenceVal.CompetenceId) {
	var competenceExporter competenceVal.CompetenceIdExporter
	val.Export(&competenceExporter)
	q.payload.CompetenceIds = append(q.payload.CompetenceIds, competenceExporter)
}
func (q *ArtifactProposedQuery) AddAuthorId(val memberVal.MemberId) {
	var authorExporter memberVal.MemberIdExporter
	val.Export(&authorExporter)
	q.payload.AuthorIds = append(q.payload.AuthorIds, authorExporter)
}
func (q *ArtifactProposedQuery) SetOwnerId(val memberVal.MemberId) {
	val.Export(&q.payload.OwnerId)
}
func (q *ArtifactProposedQuery) SetCreatedAt(val time.Time) {
	q.payload.CreatedAt = val
}

func (q *ArtifactProposedQuery) Evaluate(s infrastructure.DbSession) (infrastructure.Result, error) {
	q.EventInsertQuery.SetPayload(q.payload)
	return q.EventInsertQuery.Evaluate(s)
}

type ArtifactProposedPayload struct {
	Id            values.TenantArtifactIdExporter // Remove?
	Status        exporters.Uint8Exporter
	Name          exporters.StringExporter
	Description   exporters.StringExporter
	Url           exporters.StringExporter
	CompetenceIds []competenceVal.CompetenceIdExporter
	AuthorIds     []memberVal.MemberIdExporter
	OwnerId       memberVal.MemberIdExporter
	CreatedAt     time.Time
}
