package services

import (
	"html/template"
	"net/http"

	"github.com/sarathkumar17/go-adventure/internal/adventure"
)

func ParseStoryService(chapter adventure.Chapter, w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("ui/templates/story.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, chapter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
