package interfaces

import (
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/member/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type EndorsementExporter interface {
	SetState(
		recognizerId interfaces2.TenantMemberIdExporter,
		recognizerGrade interfaces.Exporter[uint8],
		recognizerVersion uint,
		endorsedId interfaces2.TenantMemberIdExporter,
		endorsedGrade interfaces.Exporter[uint8],
		endorsedVersion uint,
		artifactId interfaces.Exporter[uint64],
		createdAt time.Time,
	)
}
