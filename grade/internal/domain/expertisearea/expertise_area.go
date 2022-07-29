package expertisearea

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

type ExpertiseArea struct {
	id        ExpertiseAreaId
	name      Name
	createdBy member.TenantMemberId
	createdAt time.Time
}
