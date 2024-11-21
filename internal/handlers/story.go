package handlers

import (
	"net/http"
	"strings"

	"github.com/sarathkumar17/go-adventure/internal/db"
	"github.com/sarathkumar17/go-adventure/internal/services"
)

func Story(w http.ResponseWriter, r *http.Request) {
	chapterName := strings.TrimPrefix(r.URL.Path, "/story/")
	if chapterName == "" {
		http.Error(w, "Not found", http.StatusBadRequest)
		return
	}

	adventure, err := db.StartDb("gopher.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	chapter, err := services.GetChapterService(adventure, chapterName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	services.ParseStoryService(chapter, w)

}
