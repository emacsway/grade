package artifact

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/expertisearea/expertisearea"
	"time"
)

type Artifact struct {
	id               artifact.ArtifactId
	status           artifact.Status
	name             artifact.Name
	description      artifact.Description
	url              artifact.Url
	expertiseAreaIds []expertisearea.ExpertiseAreaId
	authorIds        []endorsed.EndorsedId
	createdAt        time.Time
}
