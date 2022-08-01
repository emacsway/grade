package assignment

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

func NewAssignmentFakeFactory() *AssignmentFakeFactory {
	idFactory := member.NewTenantMemberIdFakeFactory()
	idFactory.MemberId = 2
	return &AssignmentFakeFactory{
		EndorsedId:      idFactory,
		EndorsedVersion: 2,
		AssignedGrade:   1,
		Reason:          "Any",
		CreatedAt:       time.Now(),
	}
}

type AssignmentFakeFactory struct {
	EndorsedId      *member.TenantMemberIdFakeFactory
	EndorsedVersion uint
	AssignedGrade   uint8
	Reason          string
	CreatedAt       time.Time
}

func (f AssignmentFakeFactory) Create() (Assignment, error) {
	endorsedId, err := member.NewTenantMemberId(f.EndorsedId.TenantId, f.EndorsedId.MemberId)
	if err != nil {
		return Assignment{}, err
	}
	assignedGrade, err := grade.DefaultConstructor(f.AssignedGrade)
	if err != nil {
		return Assignment{}, err
	}
	reason, err := NewReason(f.Reason)
	if err != nil {
		return Assignment{}, err
	}
	return NewAssignment(endorsedId, f.EndorsedVersion, assignedGrade, reason, f.CreatedAt)
}
