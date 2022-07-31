package endorsement

import (
	"errors"
	"time"

	"github.com/hashicorp/go-multierror"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

type Weight uint8

const (
	LowerWeight  = 0
	PeerWeight   = 1
	HigherWeight = 2
)

var (
	ErrLowerGradeEndorses = errors.New(
		"it is allowed to receive endorsements only from members with equal or higher grade",
	)
	ErrEndorsementOneself = errors.New(
		"recognizer can't endorse himself",
	)
)

func CanEndorse(
	recognizerId member.TenantMemberId,
	recognizerGrade shared.Grade,
	endorsedId member.TenantMemberId,
	endorsedGrade shared.Grade,
) error {
	var err error

	if recognizerGrade.LessThan(endorsedGrade) {
		err = multierror.Append(err, ErrLowerGradeEndorses)
	}

	if recognizerId.Equals(endorsedId) {
		err = multierror.Append(err, ErrEndorsementOneself)
	}
	return err
}

func NewEndorsement(
	recognizerId member.TenantMemberId,
	recognizerGrade shared.Grade,
	recognizerVersion uint,
	endorsedId member.TenantMemberId,
	endorsedGrade shared.Grade,
	endorsedVersion uint,
	artifactId artifact.ArtifactId,
	createdAt time.Time,
) (Endorsement, error) {
	err := CanEndorse(recognizerId, recognizerGrade, endorsedId, endorsedGrade)
	if err != nil {
		return Endorsement{}, err
	}
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
	recognizerGrade   shared.Grade
	recognizerVersion uint
	endorsedId        member.TenantMemberId
	endorsedGrade     shared.Grade
	endorsedVersion   uint
	artifactId        artifact.ArtifactId
	createdAt         time.Time
}

func (e Endorsement) IsEndorsedBy(rId member.TenantMemberId, aId artifact.ArtifactId) bool {
	return e.recognizerId == rId && e.artifactId == aId
}

func (e Endorsement) GetEndorsedGrade() shared.Grade {
	return e.endorsedGrade
}

func (e Endorsement) GetWeight() Weight {
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
	SetRecognizerGrade(shared.Grade)
	SetRecognizerVersion(uint)
	SetEndorsedId(member.TenantMemberId)
	SetEndorsedGrade(shared.Grade)
	SetEndorsedVersion(uint)
	SetArtifactId(id artifact.ArtifactId)
	SetCreatedAt(time.Time)
}
