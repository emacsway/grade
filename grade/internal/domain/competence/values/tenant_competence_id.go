package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewTenantCompetenceId(tenantId, competenceInTenantId uint) (TenantCompetenceId, error) {
	tId, err := tenant.NewTenantId(tenantId)
	if err != nil {
		return TenantCompetenceId{}, err
	}
	mId, err := NewCompetenceInTenantId(competenceInTenantId)
	if err != nil {
		return TenantCompetenceId{}, err
	}
	return TenantCompetenceId{
		tenantId:             tId,
		competenceInTenantId: mId,
	}, nil
}

type TenantCompetenceId struct {
	tenantId             tenant.TenantId
	competenceInTenantId CompetenceInTenantId
}

func (cid TenantCompetenceId) TenantId() tenant.TenantId {
	return cid.tenantId
}

func (cid TenantCompetenceId) CompetenceInTenantId() CompetenceInTenantId {
	return cid.competenceInTenantId
}

func (cid TenantCompetenceId) Equal(other specification.EqualOperand) bool {
	otherTyped, ok := other.(TenantCompetenceId)
	if !ok {
		return false
	}
	return cid.tenantId.Equal(otherTyped.TenantId()) && cid.competenceInTenantId.Equal(otherTyped.CompetenceInTenantId())
}

func (cid TenantCompetenceId) Export(ex TenantCompetenceIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetCompetenceInTenantId(cid.competenceInTenantId)
}

type TenantCompetenceIdExporterSetter interface {
	SetTenantId(tenant.TenantId)
	SetCompetenceInTenantId(CompetenceInTenantId)
}
