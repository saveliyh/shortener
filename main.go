package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// create a new router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// define a route
	router.Get("/unshorten/{link}", func(w http.ResponseWriter, r *http.Request) {
		link := chi.URLParam(r, "link")
		long_link := unshorten_link(link)
		log.Println(long_link)
		w.Write([]byte(long_link))
	})
	router.Post("/shorten/{link}", func(w http.ResponseWriter, r *http.Request) {
		link := chi.URLParam(r, "link")
		short_link := get_short_link(link)
		log.Println(short_link)
		w.Write([]byte(short_link))
	})
	// start server
	log.Panic(http.ListenAndServe(":8000", router))
}
