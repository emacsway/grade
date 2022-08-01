package expertisearea

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

type ExpertiseArea struct {
	id          ExpertiseAreaId
	name        Name
	createdById member.TenantMemberId
	createdAt   time.Time
}
