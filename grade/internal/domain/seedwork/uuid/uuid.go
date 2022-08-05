package uuid

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

type Uuid = uuid.UUID

func NewUuid() Uuid {
	return Must(uuid.FromBytes(ulid.Make().Bytes()))
}

func ParseSilent(s string) Uuid {
	return Must(Parse(s))
}

func Parse(s string) (Uuid, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return Uuid{}, err
	}
	return Uuid(u), nil
}

func Must(id Uuid, err error) Uuid {
	if err != nil {
		panic(err)
	}
	return id
}
