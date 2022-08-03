package events

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
)

func NewEndorsementReceived(
	recognizerId member.TenantMemberId,
	recognizerGrade grade.Grade,
	recognizerVersion uint,
	specialistId member.TenantMemberId,
	specialistGrade grade.Grade,
	specialistVersion uint,
	artifactId artifact.TenantArtifactId,
	createdAt time.Time,
) EndorsementReceived {
	return EndorsementReceived{
		recognizerId:      recognizerId,
		recognizerGrade:   recognizerGrade,
		recognizerVersion: recognizerVersion,
		specialistId:      specialistId,
		specialistGrade:   specialistGrade,
		specialistVersion: specialistVersion,
		artifactId:        artifactId,
		createdAt:         createdAt,
	}
}

type EndorsementReceived struct {
	recognizerId      member.TenantMemberId
	recognizerGrade   grade.Grade
	recognizerVersion uint
	specialistId      member.TenantMemberId
	specialistGrade   grade.Grade
	specialistVersion uint
	artifactId        artifact.TenantArtifactId
	createdAt         time.Time
}

func (e EndorsementReceived) RecognizerId() member.TenantMemberId {
	return e.recognizerId
}

func (e EndorsementReceived) RecognizerGrade() grade.Grade {
	return e.recognizerGrade
}

func (e EndorsementReceived) RecognizerVersion() uint {
	return e.recognizerVersion
}

func (e EndorsementReceived) SpecialistId() member.TenantMemberId {
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
