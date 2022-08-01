package endorsement

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

type Weight uint8

const (
	LowerWeight  = 0
	PeerWeight   = 1
	HigherWeight = 2
)

func NewEndorsement(
	recognizerId member.TenantMemberId,
	recognizerGrade grade.Grade,
	recognizerVersion uint,
	endorsedId member.TenantMemberId,
	endorsedGrade grade.Grade,
	endorsedVersion uint,
	artifactId artifact.TenantArtifactId,
	createdAt time.Time,
) (Endorsement, error) {
	return Endorsement{
		recognizerId:      recognizerId,
		recognizerGrade:   recognizerGrade,
		recognizerVersion: recognizerVersion,
		endorsedId:        endorsedId,
		endorsedGrade:     endorsedGrade,
		endorsedVersion:   endorsedVersion,
		artifactId:        artifactId,
		createdAt:         createdAt,
	}, nil
}

type Endorsement struct {
	recognizerId      member.TenantMemberId
	recognizerGrade   grade.Grade
	recognizerVersion uint
	endorsedId        member.TenantMemberId
	endorsedGrade     grade.Grade
	endorsedVersion   uint
	artifactId        artifact.TenantArtifactId
	createdAt         time.Time
}

func (e Endorsement) IsEndorsedBy(rId member.TenantMemberId, aId artifact.TenantArtifactId) bool {
	return e.recognizerId.Equal(rId) && e.artifactId.Equal(aId)
}

func (e Endorsement) EndorsedGrade() grade.Grade {
	return e.endorsedGrade
}

func (e Endorsement) Weight() Weight {
	if e.recognizerGrade.Equal(e.endorsedGrade) {
		return PeerWeight
	} else if e.recognizerGrade.GreaterThan(e.endorsedGrade) {
		return HigherWeight
	}
	return LowerWeight
}

func (e Endorsement) Export(ex EndorsementExporterSetter) {
	ex.SetRecognizerId(e.recognizerId)
	ex.SetRecognizerGrade(e.recognizerGrade)
	ex.SetRecognizerVersion(e.recognizerVersion)
	ex.SetEndorsedId(e.endorsedId)
	ex.SetEndorsedGrade(e.endorsedGrade)
	ex.SetEndorsedVersion(e.endorsedVersion)
	ex.SetArtifactId(e.artifactId)
	ex.SetCreatedAt(e.createdAt)
}

type EndorsementExporterSetter interface {
	SetRecognizerId(member.TenantMemberId)
	SetRecognizerGrade(grade.Grade)
	SetRecognizerVersion(uint)
	SetEndorsedId(member.TenantMemberId)
	SetEndorsedGrade(grade.Grade)
	SetEndorsedVersion(uint)
	SetArtifactId(id artifact.TenantArtifactId)
	SetCreatedAt(time.Time)
}
