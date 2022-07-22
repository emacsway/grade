package expertisearea

import (
	endorsed "github.com/emacsway/qualifying-grade/grade/internal/domain/expertisearea/expertisearea"
	"time"
)

type ExpertiseArea struct {
	id        endorsed.ExpertiseAreaId
	createdAt time.Time
}
