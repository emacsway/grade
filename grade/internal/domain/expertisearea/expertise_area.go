package expertisearea

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/expertisearea/expertisearea"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

type ExpertiseArea struct {
	id        expertisearea.ExpertiseAreaId
	name      expertisearea.Name
	createdBy member.TenantMemberId
	createdAt time.Time
}
