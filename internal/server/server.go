package server

import (
	"net/http"

	"github.com/sarathkumar17/go-adventure/internal/handlers"
)

func StartServer() {

	server := http.NewServeMux()
	server.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui"))))

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/views/index.html")
	})

	server.HandleFunc("/start", handlers.Start)
	server.HandleFunc("/story/", handlers.Story)
	print("Server started on port 8080")
	http.ListenAndServe(":8080", server)
}
