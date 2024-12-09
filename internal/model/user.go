package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username       string             `json:"username,omitempty" bson:"username,omitempty"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty"`
	ShortenedLinks []UrlMap           `json:"shortened_links,omitempty" bson:"shortened_links,omitempty"`
}

type UrlMap struct {
	OriginalUrl  string `json:"original_url,omitempty" bson:"original_url,omitempty"`
	ShortenedUrl string `json:"shortened_url,omitempty" bson:"shortened_url,omitempty"`
}
