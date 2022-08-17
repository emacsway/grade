package artifact

import (
	"encoding/json"
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/competence"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type ArtifactProposedInsertQuery struct {
	params  [8]any
	payload ArtifactProposedPayload
}

func (q ArtifactProposedInsertQuery) sql() string {
	return `
		INSERT INTO recognizer
		(tenant_id, stream_type, stream_id, stream_position, event_type, event_version, payload, metadata)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)`
}
func (q *ArtifactProposedInsertQuery) SetId(val artifact.TenantArtifactId) {
	val.Export(&q.payload.Id)
	val.Export(q)
}

func (q *ArtifactProposedInsertQuery) SetTenantId(val tenant.TenantId) {
	var v exporters.UuidExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *ArtifactProposedInsertQuery) SetStreamType(val string) {
	q.params[1] = val
}

func (q *ArtifactProposedInsertQuery) SetArtifactId(val artifact.ArtifactId) {
	var v exporters.UuidExporter
	val.Export(&v)
	q.params[2] = v
}

func (q *ArtifactProposedInsertQuery) SetAggregateVersion(val uint) {
	q.params[3] = val
}

func (q *ArtifactProposedInsertQuery) SetEventType(val string) {
	q.params[4] = val
}

func (q *ArtifactProposedInsertQuery) SetEventVersion(val uint8) {
	q.params[5] = val
}

func (q *ArtifactProposedInsertQuery) SetStatus(val artifact.Status) {
	val.Export(&q.payload.Status)
}
func (q *ArtifactProposedInsertQuery) SetName(val artifact.Name) {
	val.Export(&q.payload.Name)
}
func (q *ArtifactProposedInsertQuery) SetDescription(val artifact.Description) {
	val.Export(&q.payload.Description)
}
func (q *ArtifactProposedInsertQuery) SetUrl(val artifact.Url) {
	val.Export(&q.payload.Url)
}
func (q *ArtifactProposedInsertQuery) AddCompetenceId(val competence.TenantCompetenceId) {
	var competenceExporter competence.TenantCompetenceIdExporter
	val.Export(&competenceExporter)
	q.payload.CompetenceIds = append(q.payload.CompetenceIds, competenceExporter)
}
func (q *ArtifactProposedInsertQuery) AddAuthorId(val member.TenantMemberId) {
	var authorExporter member.TenantMemberIdExporter
	val.Export(&authorExporter)
	q.payload.AuthorIds = append(q.payload.AuthorIds, authorExporter)
}
func (q *ArtifactProposedInsertQuery) SetOwnerId(val member.TenantMemberId) {
	val.Export(&q.payload.OwnerId)
}
func (q *ArtifactProposedInsertQuery) SetCreatedAt(val time.Time) {
	q.payload.CreatedAt = val
}

func (q *ArtifactProposedInsertQuery) Execute(s infrastructure.DbSession) (infrastructure.Result, error) {
	payload, err := json.Marshal(q.payload)
	if err != nil {
		return nil, err
	}
	q.params[6] = payload
	return s.Exec(q.sql(), q.params[:]...)
}

type ArtifactProposedPayload struct {
	Id            artifact.TenantArtifactIdExporter // Remove?
	Status        exporters.Uint8Exporter
	Name          exporters.StringExporter
	Description   exporters.StringExporter
	Url           exporters.StringExporter
	CompetenceIds []competence.TenantCompetenceIdExporter
	AuthorIds     []member.TenantMemberIdExporter
	OwnerId       member.TenantMemberIdExporter
	CreatedAt     time.Time
}