package handlers

import (
	"net/http"

	"github.com/sarathkumar17/go-adventure/internal/db"
	"github.com/sarathkumar17/go-adventure/internal/services"
)

func Start(w http.ResponseWriter, r *http.Request) {

	adventure, err := db.StartDb("gopher.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	chapter, err := services.StartStoryService(adventure)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	services.ParseStoryService(chapter, w)
}
