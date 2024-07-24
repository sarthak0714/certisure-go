package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username     string             `bson:"username" json:"username"`
	Email        string             `bson:"email" json:"email"`
	Password     string             `bson:"password" json:"-"`
	Phone        string             `bson:"phone" json:"phone"`
	Subscription *Subscription      `bson:"subscription" json:"subscription"`
}

type Subscription struct {
	Type       string    `bson:"type" json:"type"`
	Expiration time.Time `bson:"expiration_date" json:"expiration_date"`
}
