package queries

import (
	"fmt"
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/events"
	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/infrastructure/repository"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewArtifactProposedQuery(event *events.ArtifactProposed) *ArtifactProposedQuery {
	q := &ArtifactProposedQuery{}
	event.Export(q)
	return q
}

type ArtifactProposedQuery struct {
	repository.EventInsertQuery
	payload ArtifactProposedPayload
}

func (q *ArtifactProposedQuery) SetAggregateId(val values.ArtifactId) {
	val.Export(&q.payload.AggregateId)
	val.Export(q)
}

func (q *ArtifactProposedQuery) SetTenantId(val tenantVal.TenantId) {
	val.Export(func(v uint) { q.EventInsertQuery.SetTenantId(v) })
}

func (q *ArtifactProposedQuery) SetArtifactId(val values.InternalArtifactId) {
	val.Export(func(v uint) { q.SetStreamId(fmt.Sprintf("%d", v)) })
}

func (q *ArtifactProposedQuery) SetStatus(val values.Status) {
	val.Export(func(v uint8) { q.payload.Status = v })
}
func (q *ArtifactProposedQuery) SetName(val values.Name) {
	val.Export(func(v string) { q.payload.Name = v })
}
func (q *ArtifactProposedQuery) SetDescription(val values.Description) {
	val.Export(func(v string) { q.payload.Description = v })
}
func (q *ArtifactProposedQuery) SetUrl(val values.Url) {
	val.Export(func(v string) { q.payload.Url = v })
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

func (q *ArtifactProposedQuery) Evaluate(s session.Session) (session.Result, error) {
	q.EventInsertQuery.SetPayload(q.payload)
	return q.EventInsertQuery.Evaluate(s)
}

type ArtifactProposedPayload struct {
	AggregateId   values.ArtifactIdExporter // Remove?
	Status        uint8
	Name          string
	Url           string
	Description   string
	CompetenceIds []competenceVal.CompetenceIdExporter
	AuthorIds     []memberVal.MemberIdExporter
	OwnerId       memberVal.MemberIdExporter
	CreatedAt     time.Time
}
