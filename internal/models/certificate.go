package models

type Certificate struct {
	Id        string `bson:"_id,omitempty" json:"_id"`
	GroupId   string `bson:"group_id" json:"group_id"`
	GroupName string `bson:"group_name" json:"group_name"`
	Email     string `bson:"email" json:"email"`
	Username  string `bson:"username" json:"username"`
	Date      string `bson:"data" json:"data"`
}
