package artifact

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/expertisearea/expertisearea"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

type Artifact struct {
	id               artifact.ArtifactId
	status           artifact.Status
	name             artifact.Name
	description      artifact.Description
	url              artifact.Url
	expertiseAreaIds []expertisearea.ExpertiseAreaId
	authorIds        []member.MemberId
	createdAt        time.Time
}
