package http

import (
	"backend-app/db"
	"backend-app/model"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

// Get all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {

	collection := db.GetCollection(db.PostsCollection)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var empty interface{} = bson.M{}

	cursor, err := collection.Find(ctx, empty)
	if err != nil {
		WriteInternalServerError(w, err)
		return
	}

	var results = make([]model.Post, 0)
	for cursor.TryNext(ctx) {
		var post model.Post

		if err := cursor.Decode(&post); err != nil {
			WriteInternalServerError(w, err)
			return
		}

		results = append(results, post)
		cursor.Next(ctx)
	}

	b, err := json.Marshal(results)

	if err != nil {
		WriteInternalServerError(w, err)
	} else {
		_, _ = w.Write(b)
	}
}
