package main

import (
	"log"
	"net/http"

	"url_shortener/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Println("Start Api")
	storage, err := database.Postgres{}.Connect_db()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connect to storage")
	// create a new router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// define a route
	router.Get("/unshorten/{link}", func(w http.ResponseWriter, r *http.Request) {
		link := chi.URLParam(r, "link")
		log.Println(link)
		long_link, err := unshorten_link(link, storage)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write([]byte(long_link))
	})
	router.Post("/shorten/{link}", func(w http.ResponseWriter, r *http.Request) {
		link := chi.URLParam(r, "link")
		log.Println(link)
		short_link, err := get_short_link(link, storage)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(short_link))
	})
	// start server
	log.Panic(http.ListenAndServe(":9000", router))
}
