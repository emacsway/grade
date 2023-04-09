package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantCompetenceId(tenantId, competenceId uint) (TenantCompetenceId, error) {
	tId, err := tenant.NewTenantId(tenantId)
	if err != nil {
		return TenantCompetenceId{}, err
	}
	mId, err := NewCompetenceId(competenceId)
	if err != nil {
		return TenantCompetenceId{}, err
	}
	return TenantCompetenceId{
		tenantId:     tId,
		competenceId: mId,
	}, nil
}

type TenantCompetenceId struct {
	tenantId     tenant.TenantId
	competenceId CompetenceId
}

func (cid TenantCompetenceId) TenantId() tenant.TenantId {
	return cid.tenantId
}

func (cid TenantCompetenceId) CompetenceId() CompetenceId {
	return cid.competenceId
}

func (cid TenantCompetenceId) Equal(other specification.EqualOperand) bool {
	otherTyped, ok := other.(TenantCompetenceId)
	if !ok {
		return false
	}
	return cid.tenantId.Equal(otherTyped.TenantId()) && cid.competenceId.Equal(otherTyped.CompetenceId())
}

func (cid TenantCompetenceId) Export(ex TenantCompetenceIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetCompetenceId(cid.competenceId)
}

type TenantCompetenceIdExporterSetter interface {
	SetTenantId(tenant.TenantId)
	SetCompetenceId(CompetenceId)
}
