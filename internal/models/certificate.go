package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Certificate struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	GroupId   string             `bson:"group_id" json:"group_id"`
	GroupName string             `bson:"group_name" json:"group_name"`
	Email     string             `bson:"email" json:"email"`
	Username  string             `bson:"username" json:"username"`
	Date      time.Time          `bson:"data" json:"data"`
}
