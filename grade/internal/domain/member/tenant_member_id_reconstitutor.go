package member

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

type TenantMemberIdReconstitutor struct {
	TenantId uuid.Uuid
	MemberId uuid.Uuid
}

func (r TenantMemberIdReconstitutor) Reconstitute() (TenantMemberId, error) {
	return NewTenantMemberId(r.TenantId, r.MemberId)
}
