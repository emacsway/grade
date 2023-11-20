package events

import (
	"time"

	artifact "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

func NewEndorsementReceived(
	endorserId member.MemberId,
	endorserGrade grade.Grade,
	endorserVersion uint,
	specialistId member.MemberId,
	specialistGrade grade.Grade,
	specialistVersion uint,
	artifactId artifact.TenantArtifactId,
	createdAt time.Time,
) EndorsementReceived {
	return EndorsementReceived{
		endorserId:        endorserId,
		endorserGrade:     endorserGrade,
		endorserVersion:   endorserVersion,
		specialistId:      specialistId,
		specialistGrade:   specialistGrade,
		specialistVersion: specialistVersion,
		artifactId:        artifactId,
		createdAt:         createdAt,
	}
}

type EndorsementReceived struct {
	endorserId        member.MemberId
	endorserGrade     grade.Grade
	endorserVersion   uint
	specialistId      member.MemberId
	specialistGrade   grade.Grade
	specialistVersion uint
	artifactId        artifact.TenantArtifactId
	createdAt         time.Time
}

func (e EndorsementReceived) EndorserId() member.MemberId {
	return e.endorserId
}

func (e EndorsementReceived) EndorserGrade() grade.Grade {
	return e.endorserGrade
}

func (e EndorsementReceived) EndorserVersion() uint {
	return e.endorserVersion
}

func (e EndorsementReceived) SpecialistId() member.MemberId {
	return e.specialistId
}

func (e EndorsementReceived) SpecialistGrade() grade.Grade {
	return e.specialistGrade
}

func (e EndorsementReceived) SpecialistVersion() uint {
	return e.specialistVersion
}

func (e EndorsementReceived) ArtifactId() artifact.TenantArtifactId {
	return e.artifactId
}

func (e EndorsementReceived) CreatedAt() time.Time {
	return e.createdAt
}
