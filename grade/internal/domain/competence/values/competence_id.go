package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/specification/domain"
)

func NewCompetenceId(tenantId, competenceId uint) (CompetenceId, error) {
	tId, err := tenant.NewTenantId(tenantId)
	if err != nil {
		return CompetenceId{}, err
	}
	mId, err := NewInternalCompetenceId(competenceId)
	if err != nil {
		return CompetenceId{}, err
	}
	return CompetenceId{
		tenantId:     tId,
		competenceId: mId,
	}, nil
}

type CompetenceId struct {
	tenantId     tenant.TenantId
	competenceId InternalCompetenceId
}

func (cid CompetenceId) TenantId() tenant.TenantId {
	return cid.tenantId
}

func (cid CompetenceId) CompetenceId() InternalCompetenceId {
	return cid.competenceId
}

func (cid CompetenceId) Equal(other specification.EqualOperand) bool {
	otherTyped, ok := other.(CompetenceId)
	if !ok {
		return false
	}
	return cid.tenantId.Equal(otherTyped.TenantId()) && cid.competenceId.Equal(otherTyped.CompetenceId())
}

func (cid CompetenceId) Export(ex CompetenceIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetCompetenceId(cid.competenceId)
}

type CompetenceIdExporterSetter interface {
	SetTenantId(tenant.TenantId)
	SetCompetenceId(InternalCompetenceId)
}
