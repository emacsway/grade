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
		SpecialistId:      member.NewTenantMemberIdFakeFactory(),
		SpecialistGrade:   1,
		SpecialistVersion: 5,
		ArtifactId:        artifact.NewTenantArtifactIdFakeFactory(),
		RecognizerId:      recognizerIdFactory,
		RecognizerGrade:   2,
		RecognizerVersion: 3,
		CreatedAt:         time.Now(),
	}
}

type EndorsementFakeFactory struct {
	SpecialistId      member.TenantMemberIdFakeFactory
	SpecialistGrade   uint8
	SpecialistVersion uint
	ArtifactId        artifact.TenantArtifactIdFakeFactory
	RecognizerId      member.TenantMemberIdFakeFactory
	RecognizerGrade   uint8
	RecognizerVersion uint
	CreatedAt         time.Time
}

func (f EndorsementFakeFactory) Create() (Endorsement, error) {
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
	recognizerId, err := f.RecognizerId.Create()
	if err != nil {
		return Endorsement{}, err
	}
	recognizerGrade, err := grade.DefaultConstructor(f.RecognizerGrade)
	if err != nil {
		return Endorsement{}, err
	}
	return NewEndorsement(
		specialistId, specialistGrade, f.SpecialistVersion, artifactId,
		recognizerId, recognizerGrade, f.RecognizerVersion, f.CreatedAt,
	)
}
