package model

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Post struct {
	Id string
	Title string
	Date time.Time
	HtmlContent string
}

func PostFromBson(cursor *mongo.Cursor) (Post, error) {
	var post Post

	if err := cursor.Decode(post); err != nil {
		return post, err
	}

	return post, nil
}