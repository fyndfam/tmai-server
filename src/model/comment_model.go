package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Content   string             `json:"content" bson:"content"`
	PostID    primitive.ObjectID `json:"postId" bson:"postId"`
	CreatedBy CreatedBy          `json:"createdBy" bson:"createdBy"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
