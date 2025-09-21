package domain

import (
	"emopathy-api/shared/common_vo"
	"emopathy-api/shared/ex"
	"time"
)

var (
	ErrFcmTokenIsEmpty    = ex.NewInvalidArgument("fcm token is empty")
	ErrAccessTokenIsEmpty = ex.NewInvalidArgument("access token is empty")
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

type fcmToken struct {
	token string
}

func NewFcmToken(token string) fcmToken {
	return fcmToken{token: token}
}

func NewParseFcmToken(token string) (fcmToken, error) {
	if token == "" {
		return fcmToken{}, ErrFcmTokenIsEmpty
	}
	return fcmToken{token: token}, nil
}

func MustParseFcmToken(token string) fcmToken {
	parsedFcmToken, err := NewParseFcmToken(token)
	if err != nil {
		panic(err)
	}
	return parsedFcmToken
}

func (fcmToken fcmToken) String() string {
	return fcmToken.token
}

type accessToken struct {
	token string
}

func NewParseAccessToken(token string) (accessToken, error) {
	if token == "" {
		return accessToken{}, ErrAccessTokenIsEmpty
	}
	return accessToken{token: token}, nil
}

func MustParseAccessToken(token string) accessToken {
	parsedAccessToken, err := NewParseAccessToken(token)
	if err != nil {
		panic(err)
	}
	return parsedAccessToken
}

func (accessToken accessToken) String() string {
	return accessToken.token
}

type createdAt struct {
	time time.Time
}

func NewCreatedAt(time time.Time) createdAt {
	return createdAt{time: time}
}

func NewParseCreatedAt(time time.Time) (createdAt, error) {
	return createdAt{time: time}, nil
}

func MustParseCreatedAt(time time.Time) createdAt {
	parsedCreatedAt, err := NewParseCreatedAt(time)
	if err != nil {
		panic(err)
	}
	return parsedCreatedAt
}

func (createdAt createdAt) String() string {
	return createdAt.time.String()
}

func (createdAt createdAt) Value() time.Time {
	return createdAt.time
}

type updatedAt struct {
	time time.Time
}

func NewUpdatedAt(time time.Time) updatedAt {
	return updatedAt{time: time}
}

func NewParseUpdatedAt(time time.Time) (updatedAt, error) {
	return updatedAt{time: time}, nil
}

func MustParseUpdatedAt(time time.Time) updatedAt {
	parsedUpdatedAt, err := NewParseUpdatedAt(time)
	if err != nil {
		panic(err)
	}
	return parsedUpdatedAt
}

func (updatedAt updatedAt) String() string {
	return updatedAt.time.String()
}
