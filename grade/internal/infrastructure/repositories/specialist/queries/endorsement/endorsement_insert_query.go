package endorsement

import (
	"time"

	artifactVal "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type EndorsementInsertQuery struct {
	params [9]any
}

func (q EndorsementInsertQuery) sql() string {
	return `
		INSERT INTO endorsement (
			tenant_id, specialist_id, specialist_grade, specialist_version,
			artifact_id, endorser_id, endorser_grade, endorser_version, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT DO NOTHING`
}

func (q *EndorsementInsertQuery) SetSpecialistId(val memberVal.TenantMemberId) {
	var v memberVal.TenantMemberIdExporter
	val.Export(&v)
	q.params[0] = v.TenantId
	q.params[1] = v.MemberId
}

func (q *EndorsementInsertQuery) SetSpecialistGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.params[2] = v
}

func (q *EndorsementInsertQuery) SetSpecialistVersion(val uint) {
	q.params[3] = val
}

func (q *EndorsementInsertQuery) SetArtifactId(val artifactVal.TenantArtifactId) {
	var v artifactVal.TenantArtifactIdExporter
	val.Export(&v)
	q.params[4] = v.ArtifactId
}

func (q *EndorsementInsertQuery) SetEndorserId(val memberVal.TenantMemberId) {
	var v memberVal.TenantMemberIdExporter
	val.Export(&v)
	q.params[5] = v.MemberId
}

func (q *EndorsementInsertQuery) SetEndorserGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.params[6] = v
}

func (q *EndorsementInsertQuery) SetEndorserVersion(val uint) {
	q.params[7] = val
}

func (q *EndorsementInsertQuery) SetCreatedAt(val time.Time) {
	q.params[8] = val
}

func (q *EndorsementInsertQuery) Evaluate(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	return s.Exec(q.sql(), q.params[:]...)
}
