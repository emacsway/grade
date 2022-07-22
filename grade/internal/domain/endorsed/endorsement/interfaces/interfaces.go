package interfaces

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type EndorsementExporter interface {
	SetState(
		recognizerId interfaces.Exporter[uint64],
		recognizerGrade interfaces.Exporter[uint8],
		recognizerVersion uint,
		endorsedId interfaces.Exporter[uint64],
		endorsedGrade interfaces.Exporter[uint8],
		endorsedVersion uint,
		artifactId interfaces.Exporter[uint64],
		createdAt time.Time,
	)
}
