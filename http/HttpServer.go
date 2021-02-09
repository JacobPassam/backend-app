package http

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func Run() {

	router := chi.NewRouter()
	router.Get("/posts", GetPosts)

	http.ListenAndServe(":3000", router)
}

type InlineError struct {ErrorStr string}
func (e InlineError) Error() string {
	return e.ErrorStr
}

func WriteInternalServerError(w http.ResponseWriter, e error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	if r, err := json.Marshal(map[string]string{"error": e.Error()}); err != nil {
		_, _ = w.Write([]byte("INTERNAL SERVER ERROR"))
	} else {
		_, _ = w.Write(r)
	}
}