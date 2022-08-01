package events

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

func NewEndorsementReceived(
	recognizerId member.TenantMemberId,
	recognizerGrade grade.Grade,
	recognizerVersion uint,
	endorsedId member.TenantMemberId,
	endorsedGrade grade.Grade,
	endorsedVersion uint,
	artifactId artifact.ArtifactId,
	createdAt time.Time,
) EndorsementReceived {
	return EndorsementReceived{
		recognizerId:      recognizerId,
		recognizerGrade:   recognizerGrade,
		recognizerVersion: recognizerVersion,
		endorsedId:        endorsedId,
		endorsedGrade:     endorsedGrade,
		endorsedVersion:   endorsedVersion,
		artifactId:        artifactId,
		createdAt:         createdAt,
	}
}

type EndorsementReceived struct {
	recognizerId      member.TenantMemberId
	recognizerGrade   grade.Grade
	recognizerVersion uint
	endorsedId        member.TenantMemberId
	endorsedGrade     grade.Grade
	endorsedVersion   uint
	artifactId        artifact.ArtifactId
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

func (e EndorsementReceived) EndorsedId() member.TenantMemberId {
	return e.endorsedId
}

func (e EndorsementReceived) EndorsedGrade() grade.Grade {
	return e.endorsedGrade
}

func (e EndorsementReceived) EndorsedVersion() uint {
	return e.endorsedVersion
}

func (e EndorsementReceived) ArtifactId() artifact.ArtifactId {
	return e.artifactId
}

func (e EndorsementReceived) CreatedAt() time.Time {
	return e.createdAt
}
