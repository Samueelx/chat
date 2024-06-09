package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	hub := newHub()
	go hub.run()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		if r.URL.Path != "/" {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, "home.html")
	})

	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	sv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Printf("Server starting on port: %s\n", sv.Addr)
	err := sv.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting server: %v\n", err)
	}
}
