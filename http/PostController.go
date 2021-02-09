package http

import (
	"backend-app/db"
	"backend-app/model"
	"context"
	"encoding/json"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	collection := db.GetCollection(db.PostsCollection)

	var data interface{}
	data = json.NewDecoder(r.Body).Decode(&data)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cursor, err := collection.Find(ctx, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	results := make([]model.Post, 0)

	for cursor.TryNext(ctx) {
		post := model.Post{}
		cursor.Decode(&post)

		results = append(results, post)
		cursor.Next(ctx)
	}

	enc, _ := json.Marshal(results)
	w.Write(enc)

}
