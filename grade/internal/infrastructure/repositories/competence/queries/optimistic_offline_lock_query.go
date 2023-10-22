package queries

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type OptimisticOfflineLockLockQuery struct {
	params [5]any
}

func (q OptimisticOfflineLockLockQuery) sql() string {
	return `
		UPDATE competence
		SET version=$4
		WHERE 
			tenant_id=$1 AND competence_id=$2 AND version=$3`
}

func (q *OptimisticOfflineLockLockQuery) SetId(val values.TenantCompetenceId) {
	val.Export(q)
}

func (q *OptimisticOfflineLockLockQuery) SetTenantId(val tenantVal.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *OptimisticOfflineLockLockQuery) SetCompetenceInTenantId(val values.CompetenceInTenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *OptimisticOfflineLockLockQuery) SetName(val values.Name) {
}

func (q *OptimisticOfflineLockLockQuery) SetOwnerId(val memberVal.TenantMemberId) {
}

func (q *OptimisticOfflineLockLockQuery) SetCreatedAt(val time.Time) {
}

func (q *OptimisticOfflineLockLockQuery) SetInitialVersion(val uint) {
	q.params[2] = val
}

func (q *OptimisticOfflineLockLockQuery) SetVersion(val uint) {
	q.params[3] = val
}

func (q *OptimisticOfflineLockLockQuery) Evaluate(s infrastructure.DbSession) (infrastructure.Result, error) {
	result, err := s.Exec(q.sql(), q.params[:]...)
	if err != nil {
		return result, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return result, err
	}
	if rowsAffected == 0 {
		return result, aggregate.ErrConcurrency
	}
	return result, err
}
