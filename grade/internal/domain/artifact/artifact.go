package artifact

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/expertisearea/expertisearea"
)

type Artifact struct {
	id               artifact.ArtifactId
	expertiseAreaIds []expertisearea.ExpertiseAreaId
	authors          []endorsed.EndorsedId
}
