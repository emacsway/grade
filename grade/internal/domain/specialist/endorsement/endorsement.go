package endorsement

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
)

type Weight uint8

const (
	LowerWeight  = 0
	PeerWeight   = 1
	HigherWeight = 2
)

func NewEndorsement(
	specialistId member.TenantMemberId,
	specialistGrade grade.Grade,
	specialistVersion uint,
	artifactId artifact.TenantArtifactId,
	recognizerId member.TenantMemberId,
	recognizerGrade grade.Grade,
	recognizerVersion uint,
	createdAt time.Time,
) (Endorsement, error) {
	return Endorsement{
		specialistId:      specialistId,
		specialistGrade:   specialistGrade,
		specialistVersion: specialistVersion,
		artifactId:        artifactId,
		recognizerId:      recognizerId,
		recognizerGrade:   recognizerGrade,
		recognizerVersion: recognizerVersion,
		createdAt:         createdAt,
	}, nil
}

type Endorsement struct {
	specialistId      member.TenantMemberId
	specialistGrade   grade.Grade
	specialistVersion uint
	artifactId        artifact.TenantArtifactId
	recognizerId      member.TenantMemberId
	recognizerGrade   grade.Grade
	recognizerVersion uint
	createdAt         time.Time
}

func (e Endorsement) IsEndorsedBy(rId member.TenantMemberId, aId artifact.TenantArtifactId) bool {
	return e.recognizerId.Equal(rId) && e.artifactId.Equal(aId)
}

func (e Endorsement) SpecialistGrade() grade.Grade {
	return e.specialistGrade
}

func (e Endorsement) Weight() Weight {
	if e.recognizerGrade.Equal(e.specialistGrade) {
		return PeerWeight
	} else if e.recognizerGrade.GreaterThan(e.specialistGrade) {
		return HigherWeight
	}
	return LowerWeight
}

func (e Endorsement) Export(ex EndorsementExporterSetter) {
	ex.SetSpecialistId(e.specialistId)
	ex.SetSpecialistGrade(e.specialistGrade)
	ex.SetSpecialistVersion(e.specialistVersion)
	ex.SetArtifactId(e.artifactId)
	ex.SetRecognizerId(e.recognizerId)
	ex.SetRecognizerGrade(e.recognizerGrade)
	ex.SetRecognizerVersion(e.recognizerVersion)
	ex.SetCreatedAt(e.createdAt)
}

type EndorsementExporterSetter interface {
	SetSpecialistId(member.TenantMemberId)
	SetSpecialistGrade(grade.Grade)
	SetSpecialistVersion(uint)
	SetArtifactId(artifact.TenantArtifactId)
	SetRecognizerId(member.TenantMemberId)
	SetRecognizerGrade(grade.Grade)
	SetRecognizerVersion(uint)
	SetCreatedAt(time.Time)
}
