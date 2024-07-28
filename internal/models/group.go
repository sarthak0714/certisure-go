package models

type Group struct {
	Id        string  `bson:"_id,omitempty" json:"_id"`
	Date      string  `bson:"date" json:"date"`
	GroupName string  `bson:"group_name" json:"group_name"`
	Type      int     `bson:"type" json:"type"`
	ImageLink string  `bson:"imagelink" json:"imagelink"`
	Issuer    string  `bson:"issuer" json:"issuer"`
	GroupID   string  `bson:"group_id" json:"group_id"`
	FontName  string  `bson:"fontname" json:"fontname"`
	NameX     float64 `bson:"name_x" json:"name_x"`
	NameY     float64 `bson:"name_y" json:"name_y"`
	LinkX     float64 `bson:"link_x" json:"link_x"`
	LinkY     float64 `bson:"link_y" json:"link_y"`
	QrX       float64 `bson:"qr_x" json:"qr_x"`
	QrY       float64 `bson:"qr_y" json:"qr_y"`
	Email     string  `bson:"email" json:"email"`
}
