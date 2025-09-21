package domain

import "emopathy-api/shared/common_vo"

type RootEntity struct {
	id          ID
	fcmToken    fcmToken
	accessToken accessToken
	createdAt   createdAt
	updatedAt   updatedAt
}

func NewRootEntity(fcmToken fcmToken, accessToken accessToken, clock common_vo.Clock) *RootEntity {
	return &RootEntity{
		id:          RandomID(),
		fcmToken:    fcmToken,
		accessToken: accessToken,
		createdAt:   NewCreatedAt(clock.Now()),
		updatedAt:   NewUpdatedAt(clock.Now()),
	}
}
