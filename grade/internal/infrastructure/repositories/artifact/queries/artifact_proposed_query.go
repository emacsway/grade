package queries

import (
	"encoding/json"
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type ArtifactProposedQuery struct {
	params  [8]any
	payload ArtifactProposedPayload
	meta    aggregate.EventMetaExporter
}

func (q ArtifactProposedQuery) sql() string {
	return `
		INSERT INTO event_log
		(tenant_id, stream_type, stream_id, stream_position, event_type, event_version, payload, metadata)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)`
}
func (q *ArtifactProposedQuery) SetId(val values.TenantArtifactId) {
	val.Export(&q.payload.Id)
	val.Export(q)
}

func (q *ArtifactProposedQuery) SetTenantId(val tenantVal.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *ArtifactProposedQuery) SetStreamType(val string) {
	q.params[1] = val
}

func (q *ArtifactProposedQuery) SetArtifactId(val values.ArtifactId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[2] = v.String()
}

func (q *ArtifactProposedQuery) SetAggregateVersion(val uint) {
	q.params[3] = val
}

func (q *ArtifactProposedQuery) SetEventType(val string) {
	q.params[4] = val
}

func (q *ArtifactProposedQuery) SetEventMeta(val aggregate.EventMeta) {
	val.Export(&q.meta)
}

func (q *ArtifactProposedQuery) SetEventVersion(val uint8) {
	q.params[5] = val
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
func (q *ArtifactProposedQuery) AddCompetenceId(val competenceVal.TenantCompetenceId) {
	var competenceExporter competenceVal.TenantCompetenceIdExporter
	val.Export(&competenceExporter)
	q.payload.CompetenceIds = append(q.payload.CompetenceIds, competenceExporter)
}
func (q *ArtifactProposedQuery) AddAuthorId(val memberVal.TenantMemberId) {
	var authorExporter memberVal.TenantMemberIdExporter
	val.Export(&authorExporter)
	q.payload.AuthorIds = append(q.payload.AuthorIds, authorExporter)
}
func (q *ArtifactProposedQuery) SetOwnerId(val memberVal.TenantMemberId) {
	val.Export(&q.payload.OwnerId)
}
func (q *ArtifactProposedQuery) SetCreatedAt(val time.Time) {
	q.payload.CreatedAt = val
}

func (q *ArtifactProposedQuery) Evaluate(s infrastructure.DbSession) (infrastructure.Result, error) {
	payload, err := json.Marshal(q.payload)
	if err != nil {
		return nil, err
	}
	q.params[6] = payload
	meta, err := json.Marshal(q.meta)
	if err != nil {
		return nil, err
	}
	q.params[7] = meta
	return s.Exec(q.sql(), q.params[:]...)
}

type ArtifactProposedPayload struct {
	Id            values.TenantArtifactIdExporter // Remove?
	Status        exporters.Uint8Exporter
	Name          exporters.StringExporter
	Description   exporters.StringExporter
	Url           exporters.StringExporter
	CompetenceIds []competenceVal.TenantCompetenceIdExporter
	AuthorIds     []memberVal.TenantMemberIdExporter
	OwnerId       memberVal.TenantMemberIdExporter
	CreatedAt     time.Time
}
