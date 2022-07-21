package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEndorsedCreateMemento(t *testing.T) {
	f := NewEndorsedFakeFactory()
	agg, _ := f.Create()
	assert.Equal(t, f.CreateMemento(), agg.CreateMemento())
}

func NewEndorsedFakeFactory() *EndorsedFakeFactory {
	return &EndorsedFakeFactory{
		1, 2, 0, []endorsement.EndorsementFakeFactory{}, 1, time.Now(),
	}
}

type EndorsedFakeFactory struct {
	Id                   uint64
	UserId               uint64
	Grade                uint8
	ReceivedEndorsements []endorsement.EndorsementFakeFactory
	Version              uint
	CreatedAt            time.Time
}

func (f EndorsedFakeFactory) Create() (*Endorsed, error) {
	var receivedEndorsements []endorsement.Endorsement
	for i, v := range f.ReceivedEndorsements {
		e, err := v.Create()
		if err != nil {
			return nil, err
		}
		receivedEndorsements[i] = e
	}
	id, _ := endorsed.NewEndorsedId(f.Id)
	userId, _ := external.NewUserId(f.UserId)
	grade, _ := shared.NewGrade(f.Grade)
	return NewEndorsed(id, userId, grade, receivedEndorsements, f.Version, f.CreatedAt)
}

func (f EndorsedFakeFactory) CreateMemento() EndorsedMemento {
	var receivedEndorsements []endorsement.EndorsementMemento
	for i, v := range f.ReceivedEndorsements {
		receivedEndorsements[i] = v.CreateMemento()
	}
	return EndorsedMemento{
		f.Id, f.UserId, f.Grade, receivedEndorsements, f.Version, f.CreatedAt,
	}
}
