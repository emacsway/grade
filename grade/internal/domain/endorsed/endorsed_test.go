package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"testing"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/stretchr/testify/assert"
)

func TestEndorsedExport(t *testing.T) {
	ef, err := NewEndorsedFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	rf, err := recognizer.NewRecognizerFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for i := 0; i < 4; i++ {
		err := ef.ReceiveEndorsement(rf)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
	agg, err := ef.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, EndorsedState{
		Id:    ef.Id,
		Grade: ef.Grade + 1,
		ReceivedEndorsements: []endorsement.EndorsementState{
			{
				RecognizerId:      rf.Id,
				RecognizerGrade:   rf.Grade,
				RecognizerVersion: 0,
				EndorsedId:        ef.Id,
				EndorsedGrade:     ef.Grade,
				EndorsedVersion:   0,
				ArtifactId:        ef.ReceivedEndorsements[0].ArtifactId,
				CreatedAt:         ef.ReceivedEndorsements[0].CreatedAt,
			},
			{
				RecognizerId:      rf.Id,
				RecognizerGrade:   rf.Grade,
				RecognizerVersion: 0,
				EndorsedId:        ef.Id,
				EndorsedGrade:     ef.Grade,
				EndorsedVersion:   1,
				ArtifactId:        ef.ReceivedEndorsements[1].ArtifactId,
				CreatedAt:         ef.ReceivedEndorsements[1].CreatedAt,
			},
			{
				RecognizerId:      rf.Id,
				RecognizerGrade:   rf.Grade,
				RecognizerVersion: 0,
				EndorsedId:        ef.Id,
				EndorsedGrade:     ef.Grade,
				EndorsedVersion:   2,
				ArtifactId:        ef.ReceivedEndorsements[2].ArtifactId,
				CreatedAt:         ef.ReceivedEndorsements[2].CreatedAt,
			},
			{
				RecognizerId:      rf.Id,
				RecognizerGrade:   rf.Grade,
				RecognizerVersion: 0,
				EndorsedId:        ef.Id,
				EndorsedGrade:     ef.Grade + 1,
				EndorsedVersion:   3,
				ArtifactId:        ef.ReceivedEndorsements[3].ArtifactId,
				CreatedAt:         ef.ReceivedEndorsements[3].CreatedAt,
			},
		},
		GradeLogEntries: []gradelogentry.GradeLogEntryState{
			{
				EndorsedId:      ef.Id,
				EndorsedVersion: 2,
				AssignedGrade:   ef.Grade + 1,
				Reason:          "Endorsement count is achieved",
				CreatedAt:       ef.ReceivedEndorsements[2].CreatedAt,
			},
		},
		Version:   4,
		CreatedAt: ef.CreatedAt,
	}, agg.Export())
}

func TestEndorsedExportTo(t *testing.T) {
	var actualExporter EndorsedExporter
	ef, err := NewEndorsedFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	rf, err := recognizer.NewRecognizerFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for i := 0; i < 4; i++ {
		err := ef.ReceiveEndorsement(rf)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
	agg, err := ef.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.ExportTo(&actualExporter)
	assert.Equal(t, EndorsedExporter{
		Id:    seedwork.NewUint64Exporter(ef.Id),
		Grade: seedwork.NewUint8Exporter(ef.Grade + 1),
		ReceivedEndorsements: []interfaces2.EndorsementExporter{
			&endorsement.EndorsementExporter{
				RecognizerId:      seedwork.NewUint64Exporter(rf.Id),
				RecognizerGrade:   seedwork.NewUint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        seedwork.NewUint64Exporter(ef.Id),
				EndorsedGrade:     seedwork.NewUint8Exporter(ef.Grade),
				EndorsedVersion:   0,
				ArtifactId:        seedwork.NewUint64Exporter(ef.ReceivedEndorsements[0].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[0].CreatedAt,
			},
			&endorsement.EndorsementExporter{
				RecognizerId:      seedwork.NewUint64Exporter(rf.Id),
				RecognizerGrade:   seedwork.NewUint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        seedwork.NewUint64Exporter(ef.Id),
				EndorsedGrade:     seedwork.NewUint8Exporter(ef.Grade),
				EndorsedVersion:   1,
				ArtifactId:        seedwork.NewUint64Exporter(ef.ReceivedEndorsements[1].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[1].CreatedAt,
			},
			&endorsement.EndorsementExporter{
				RecognizerId:      seedwork.NewUint64Exporter(rf.Id),
				RecognizerGrade:   seedwork.NewUint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        seedwork.NewUint64Exporter(ef.Id),
				EndorsedGrade:     seedwork.NewUint8Exporter(ef.Grade),
				EndorsedVersion:   2,
				ArtifactId:        seedwork.NewUint64Exporter(ef.ReceivedEndorsements[2].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[2].CreatedAt,
			},
			&endorsement.EndorsementExporter{
				RecognizerId:      seedwork.NewUint64Exporter(rf.Id),
				RecognizerGrade:   seedwork.NewUint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        seedwork.NewUint64Exporter(ef.Id),
				EndorsedGrade:     seedwork.NewUint8Exporter(ef.Grade + 1),
				EndorsedVersion:   3,
				ArtifactId:        seedwork.NewUint64Exporter(ef.ReceivedEndorsements[3].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[3].CreatedAt,
			},
		},
		GradeLogEntries: []interfaces.GradeLogEntryExporter{
			&gradelogentry.GradeLogEntryExporter{
				EndorsedId:      seedwork.NewUint64Exporter(ef.Id),
				EndorsedVersion: 2,
				AssignedGrade:   seedwork.NewUint8Exporter(ef.Grade + 1),
				Reason:          seedwork.NewStringExporter("Endorsement count is achieved"),
				CreatedAt:       ef.ReceivedEndorsements[2].CreatedAt,
			},
		},
		Version:   4,
		CreatedAt: ef.CreatedAt,
	}, actualExporter)
}
