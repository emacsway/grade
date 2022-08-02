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
	specialistId member.TenantMemberId,
	specialistGrade grade.Grade,
	specialistVersion uint,
	artifactId artifact.TenantArtifactId,
	createdAt time.Time,
) (Endorsement, error) {
	return Endorsement{
		recognizerId:      recognizerId,
		recognizerGrade:   recognizerGrade,
		recognizerVersion: recognizerVersion,
		specialistId:      specialistId,
		specialistGrade:   specialistGrade,
		specialistVersion: specialistVersion,
		artifactId:        artifactId,
		createdAt:         createdAt,
	}, nil
}

type Endorsement struct {
	recognizerId      member.TenantMemberId
	recognizerGrade   grade.Grade
	recognizerVersion uint
	specialistId      member.TenantMemberId
	specialistGrade   grade.Grade
	specialistVersion uint
	artifactId        artifact.TenantArtifactId
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
	ex.SetRecognizerId(e.recognizerId)
	ex.SetRecognizerGrade(e.recognizerGrade)
	ex.SetRecognizerVersion(e.recognizerVersion)
	ex.SetSpecialistId(e.specialistId)
	ex.SetSpecialistGrade(e.specialistGrade)
	ex.SetSpecialistVersion(e.specialistVersion)
	ex.SetArtifactId(e.artifactId)
	ex.SetCreatedAt(e.createdAt)
}

type EndorsementExporterSetter interface {
	SetRecognizerId(member.TenantMemberId)
	SetRecognizerGrade(grade.Grade)
	SetRecognizerVersion(uint)
	SetSpecialistId(member.TenantMemberId)
	SetSpecialistGrade(grade.Grade)
	SetSpecialistVersion(uint)
	SetArtifactId(id artifact.TenantArtifactId)
	SetCreatedAt(time.Time)
}
