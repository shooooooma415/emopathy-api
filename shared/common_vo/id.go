package common_vo

import (
	"emopathy-api/shared/ex"

	"github.com/google/uuid"
)

var (
	ErrInvalidIDFormat = ex.NewInvalidArgument("invalid id format")
	ErrIDIsEmpty       = ex.NewInvalidArgument("id is empty")
)

type ID struct {
	ID uuid.UUID
}

func NewID(value uuid.UUID) (ID, error) {
	id := ID{ID: value}
	if id.ID == uuid.Nil {
		return ID{}, ErrIDIsEmpty
	}
	return id, nil
}

func MustID(value uuid.UUID) ID {
	id, err := NewID(value)
	if err != nil {
		panic(ex.Wrap(err))
	}
	return id
}

func RandomID() ID {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(ex.Wrap(err))
	}
	return MustID(id)
}

func ParseID(value string) (ID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return ID{}, ex.Wrap(ErrInvalidIDFormat, err.Error())
	}
	return NewID(id)
}

func MustParseID(value string) ID {
	id, err := ParseID(value)
	if err != nil {
		panic(ex.Wrap(err))
	}
	return id
}

func (id ID) String() string {
	return id.ID.String()
}

func (id ID) Value() uuid.UUID {
	return id.ID
}
