package interfaces

import (
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/member/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type GradeLogEntryExporter interface {
	SetState(
		endorsedId interfaces2.TenantMemberIdExporter,
		endorsedVersion uint,
		assignedGrade interfaces.Exporter[uint8],
		reason interfaces.Exporter[string],
		createdAt time.Time,
	)
}
