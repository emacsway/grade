package seedwork

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

type Uuid = uuid.UUID

func NewUuid() Uuid {
	u, err := uuid.FromBytes(ulid.Make().Bytes())
	if err != nil {
		panic(err)
	}
	return Uuid(u)
}

func Parse(s string) (Uuid, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return Uuid{}, err
	}
	return Uuid(u), nil
}
