package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewEndorsedFakeFactory() *EndorsedFakeFactory {
	return &EndorsedFakeFactory{
		1, 2, 0, []*endorsement.EndorsementFakeFactory{}, 1, time.Now(), 5,
	}
}

type EndorsedFakeFactory struct {
	Id                   uint64
	MemberId             uint64
	Grade                uint8
	ReceivedEndorsements []*endorsement.EndorsementFakeFactory
	Version              uint
	CreatedAt            time.Time
	CurrentArtifactId    uint64
}

func (f *EndorsedFakeFactory) AddReceivedEndorsement(r *recognizer.RecognizerFakeFactory) {
	e := endorsement.NewEndorsementFakeFactory()
	e.EndorsedId = f.Id
	e.EndorsedGrade = f.Grade
	e.EndorsedVersion = f.Version
	e.RecognizerId = r.Id
	e.RecognizerGrade = r.Grade
	e.RecognizerVersion = r.Version
	e.ArtifactId = f.CurrentArtifactId
	e.CreatedAt = time.Now()
	f.CurrentArtifactId += 1
	f.ReceivedEndorsements = append(f.ReceivedEndorsements, e)
}

func (f EndorsedFakeFactory) Create() (*Endorsed, error) {
	var receivedEndorsements []endorsement.Endorsement
	for _, v := range f.ReceivedEndorsements {
		e, err := v.Create()
		if err != nil {
			return nil, err
		}
		receivedEndorsements = append(receivedEndorsements, e)
	}
	id, _ := endorsed.NewEndorsedId(f.Id)
	memberId, _ := external.NewMemberId(f.MemberId)
	grade, _ := shared.NewGrade(f.Grade)
	return NewEndorsed(id, memberId, grade, receivedEndorsements, f.Version, f.CreatedAt)
}

func (f EndorsedFakeFactory) Export() EndorsedState {
	var receivedEndorsements []endorsement.EndorsementState
	for _, v := range f.ReceivedEndorsements {
		receivedEndorsements = append(receivedEndorsements, v.Export())
	}
	return EndorsedState{
		f.Id, f.MemberId, f.Grade, receivedEndorsements, f.Version, f.CreatedAt,
	}
}

func (f EndorsedFakeFactory) ExportTo(ex interfaces.EndorsedExporter) {
	var id, memberId seedwork.Uint64Exporter
	var grade seedwork.Uint8Exporter
	var receivedEndorsements []interfaces2.EndorsementExporter

	for _, v := range f.ReceivedEndorsements {
		re := &endorsement.EndorsementExporter{}
		v.ExportTo(re)
		receivedEndorsements = append(receivedEndorsements, re)
	}

	id.SetState(f.Id)
	memberId.SetState(f.MemberId)
	grade.SetState(f.Grade)
	ex.SetState(
		&id, &memberId, &grade, receivedEndorsements, f.Version, f.CreatedAt,
	)
}
