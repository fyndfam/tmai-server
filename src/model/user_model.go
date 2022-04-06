package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Email          string             `json:"email" bson:"email" binding:"required"`
	ExternalUserId string             `json:"externalUserId" bson:"externalUserId"`
	Username       *string            `json:"username,omitempty" bson:"username,omitempty"`
	Avatar         *string            `json:"avatar,omitempty" bson:"avatar,omitempty"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt" bson:"updatedAt"`
}
