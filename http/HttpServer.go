package http

import (
	"github.com/go-chi/chi"
	"net/http"
)

func Run() {

	router := chi.NewRouter()
	router.Get("/posts", GetPosts)

	http.ListenAndServe(":3000", router)
}
