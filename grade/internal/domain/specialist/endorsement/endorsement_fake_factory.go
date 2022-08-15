package endorsement

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/recognizer"
)

func NewEndorsementFakeFactory() EndorsementFakeFactory {
	recognizerIdFactory := member.NewTenantMemberIdFakeFactory()
	recognizerIdFactory.MemberId = recognizer.RecognizerMemberIdFakeValue
	return EndorsementFakeFactory{
		RecognizerId:      recognizerIdFactory,
		RecognizerGrade:   2,
		RecognizerVersion: 3,
		SpecialistId:      member.NewTenantMemberIdFakeFactory(),
		SpecialistGrade:   1,
		SpecialistVersion: 5,
		ArtifactId:        artifact.NewTenantArtifactIdFakeFactory(),
		CreatedAt:         time.Now(),
	}
}

type EndorsementFakeFactory struct {
	RecognizerId      member.TenantMemberIdFakeFactory
	RecognizerGrade   uint8
	RecognizerVersion uint
	SpecialistId      member.TenantMemberIdFakeFactory
	SpecialistGrade   uint8
	SpecialistVersion uint
	ArtifactId        artifact.TenantArtifactIdFakeFactory
	CreatedAt         time.Time
}

func (f EndorsementFakeFactory) Create() (Endorsement, error) {
	recognizerId, err := f.RecognizerId.Create()
	if err != nil {
		return Endorsement{}, err
	}
	recognizerGrade, err := grade.DefaultConstructor(f.RecognizerGrade)
	if err != nil {
		return Endorsement{}, err
	}
	specialistId, err := f.SpecialistId.Create()
	if err != nil {
		return Endorsement{}, err
	}
	specialistGrade, err := grade.DefaultConstructor(f.SpecialistGrade)
	if err != nil {
		return Endorsement{}, err
	}
	artifactId, err := f.ArtifactId.Create()
	if err != nil {
		return Endorsement{}, err
	}
	return NewEndorsement(
		recognizerId, recognizerGrade, f.RecognizerVersion,
		specialistId, specialistGrade, f.SpecialistVersion,
		artifactId, f.CreatedAt,
	)
}
