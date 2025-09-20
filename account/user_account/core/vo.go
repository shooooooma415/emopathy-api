package core

import (
	"emopathy-api/shared/common_vo"
	"emopathy-api/shared/ex"
)

type ID struct {
	id common_vo.ID
}

func RandomID() ID {
	return ID{id: common_vo.RandomID()}
}	

func NewParseID(id string) (ID, error) {
	parsedID, err := common_vo.ParseID(id)
	if err != nil {
		return ID{}, ex.Wrap(err)
	}
	return ID{id: parsedID}, nil
}

func MustParseID(id string) ID {
	return ID{id: common_vo.MustParseID(id)}
}

func (id ID) String() string {
	return id.id.String()
}

func (id ID) Value() common_vo.ID {
	return id.id
}