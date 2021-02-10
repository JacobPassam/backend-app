package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string
	Date        time.Time
	HtmlContent string
}
