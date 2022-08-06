package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member"
)

func NewCompetence(
	id TenantCompetenceId,
	name Name,
	ownerId member.TenantMemberId,
	createdAt time.Time,
) (*Competence, error) {
	return &Competence{
		id:        id,
		name:      name,
		ownerId:   ownerId,
		createdAt: createdAt,
	}, nil
}

type Competence struct {
	id        TenantCompetenceId
	name      Name
	ownerId   member.TenantMemberId
	createdAt time.Time
}
