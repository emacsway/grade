package gradelogentry

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

func NewGradeLogEntryFakeFactory() *GradeLogEntryFakeFactory {
	idFactory := member.NewTenantMemberIdFakeFactory()
	idFactory.MemberId = 2
	return &GradeLogEntryFakeFactory{
		EndorsedId:      idFactory,
		EndorsedVersion: 2,
		AssignedGrade:   1,
		Reason:          "Any",
		CreatedAt:       time.Now(),
	}
}

type GradeLogEntryFakeFactory struct {
	EndorsedId      *member.TenantMemberIdFakeFactory
	EndorsedVersion uint
	AssignedGrade   uint8
	Reason          string
	CreatedAt       time.Time
}

func (f GradeLogEntryFakeFactory) Create() (GradeLogEntry, error) {
	endorsedId, err := member.NewTenantMemberId(f.EndorsedId.TenantId, f.EndorsedId.MemberId)
	if err != nil {
		return GradeLogEntry{}, err
	}
	assignedGrade, err := grade.DefaultConstructor(f.AssignedGrade)
	if err != nil {
		return GradeLogEntry{}, err
	}
	reason, err := NewReason(f.Reason)
	if err != nil {
		return GradeLogEntry{}, err
	}
	return NewGradeLogEntry(endorsedId, f.EndorsedVersion, assignedGrade, reason, f.CreatedAt)
}
