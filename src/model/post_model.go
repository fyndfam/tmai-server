package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatedBy struct {
	Username string `json:"username" bson:"username"`
	Avatar   string `json:"avatar" bson:"avatar"`
}

type PostModel struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Content       string             `json:"content" bson:"content"`
	Tags          []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	View          int64              `json:"view" bson:"view"`
	CreatedBy     CreatedBy          `json:"createdBy" bson:"createdBy"`
	ContentEdited bool               `json:"contentEdited" bson:"contentEdited"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
	ReplyTo       primitive.ObjectID `json:"replyTo,omitempty" bson:"replyTo,omitempty"`
}
