package gradelogentry

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewGradeLogEntryFakeFactory() (*GradeLogEntryFakeFactory, error) {
	return &GradeLogEntryFakeFactory{
		EndorsedId:      1,
		EndorsedVersion: 2,
		AssignedGrade:   1,
		Reason:          "Any",
		CreatedAt:       time.Now(),
	}, nil
}

type GradeLogEntryFakeFactory struct {
	EndorsedId      uint64
	EndorsedVersion uint
	AssignedGrade   uint8
	Reason          string
	CreatedAt       time.Time
}

func (f GradeLogEntryFakeFactory) Create() (GradeLogEntry, error) {
	endorsedId, err := member.NewMemberId(f.EndorsedId)
	if err != nil {
		return GradeLogEntry{}, err
	}
	assignedGrade, err := shared.NewGrade(f.AssignedGrade)
	if err != nil {
		return GradeLogEntry{}, err
	}
	reason, err := gradelogentry.NewReason(f.Reason)
	if err != nil {
		return GradeLogEntry{}, err
	}
	return NewGradeLogEntry(endorsedId, f.EndorsedVersion, assignedGrade, reason, f.CreatedAt)
}
