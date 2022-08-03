package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member"
)

type Competence struct {
	id        CompetenceId
	name      Name
	ownerId   member.TenantMemberId
	createdAt time.Time
}
