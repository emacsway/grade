package interfaces

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

type EndorsementExporter interface {
	SetState(
		recognizerId seedwork.Uint64Exporter,
		recognizerGrade seedwork.Uint8Exporter,
		recognizerVersion uint,
		endorsedId seedwork.Uint64Exporter,
		endorsedGrade seedwork.Uint8Exporter,
		endorsedVersion uint,
		artifactId seedwork.Uint64Exporter,
		createdAt time.Time,
	)
}
