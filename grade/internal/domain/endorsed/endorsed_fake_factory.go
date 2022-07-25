package endorsed

import (
	interfaces3 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/interfaces"
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewEndorsedFakeFactory() *EndorsedFakeFactory {
	return &EndorsedFakeFactory{
		1, 2, 0, []*endorsement.EndorsementFakeFactory{},
		[]*gradelogentry.GradeLogEntryFakeFactory{}, 1, time.Now(), 5,
	}
}

type EndorsedFakeFactory struct {
	Id                   uint64
	MemberId             uint64
	Grade                uint8
	ReceivedEndorsements []*endorsement.EndorsementFakeFactory
	GradeLogEntries      []*gradelogentry.GradeLogEntryFakeFactory
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
	var gradeLogEntries []gradelogentry.GradeLogEntry
	for _, v := range f.ReceivedEndorsements {
		e, err := v.Create()
		if err != nil {
			return nil, err
		}
		receivedEndorsements = append(receivedEndorsements, e)
	}
	for _, v := range f.GradeLogEntries {
		gle, err := v.Create()
		if err != nil {
			return nil, err
		}
		gradeLogEntries = append(gradeLogEntries, gle)
	}
	id, _ := endorsed.NewEndorsedId(f.Id)
	memberId, _ := external.NewMemberId(f.MemberId)
	grade, _ := shared.NewGrade(f.Grade)
	return NewEndorsed(id, memberId, grade, receivedEndorsements, gradeLogEntries, f.Version, f.CreatedAt)
}

func (f EndorsedFakeFactory) Export() EndorsedState {
	var receivedEndorsements []endorsement.EndorsementState
	var gradeLogEntries []gradelogentry.GradeLogEntryState
	for _, v := range f.ReceivedEndorsements {
		receivedEndorsements = append(receivedEndorsements, v.Export())
	}
	for _, v := range f.GradeLogEntries {
		gradeLogEntries = append(gradeLogEntries, v.Export())
	}
	return EndorsedState{
		f.Id, f.MemberId, f.Grade, receivedEndorsements,
		gradeLogEntries, f.Version, f.CreatedAt,
	}
}

func (f EndorsedFakeFactory) ExportTo(ex interfaces.EndorsedExporter) {
	var id, memberId seedwork.Uint64Exporter
	var grade seedwork.Uint8Exporter
	var receivedEndorsements []interfaces2.EndorsementExporter
	var gradeLogEntries []interfaces3.GradeLogEntryExporter

	for _, v := range f.ReceivedEndorsements {
		re := &endorsement.EndorsementExporter{}
		v.ExportTo(re)
		receivedEndorsements = append(receivedEndorsements, re)
	}

	for _, v := range f.GradeLogEntries {
		gle := &gradelogentry.GradeLogEntryExporter{}
		v.ExportTo(gle)
		gradeLogEntries = append(gradeLogEntries, gle)
	}

	id.SetState(f.Id)
	memberId.SetState(f.MemberId)
	grade.SetState(f.Grade)
	ex.SetState(
		&id, &memberId, &grade, receivedEndorsements, gradeLogEntries, f.Version, f.CreatedAt,
	)
}
