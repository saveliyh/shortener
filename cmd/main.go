package main

import (
	"log"
	"net/http"
	"os"

	"url_shortener/internal/database"
	"url_shortener/internal/logic"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Println("Start Api")
	var (
		storage database.Storage
		err     error
	)

	args := os.Args
	if len(args) < 2 {
		log.Println("Store data in memmory(default)")
		storage, err = (&database.InMemory{}).Connect_db()
	} else if args[1] == "postgres" {
		log.Println("Store data in postgres")
		storage, err = (&database.Postgres{}).Connect_db()
		if postgres, ok := storage.(*database.Postgres); ok {
			defer postgres.Close()
		}

	} else if args[1] == "memmory" {
		storage, err = (&database.InMemory{}).Connect_db()
		log.Println("Store data in memmory")
	} else {
		log.Panic(args[1] + " is not correct storage type use \"--postgres\" or \"--memmory\"")
	}

	if err != nil {
		log.Panic(err)
	}

	// create a new router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// define a route
	router.Get("/{link}", func(w http.ResponseWriter, r *http.Request) {
		link := chi.URLParam(r, "link")
		log.Println(link)
		long_link, err := logic.Unshorten_link(link, storage)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write([]byte(long_link))
	})
	router.Post("/shorten", func(w http.ResponseWriter, r *http.Request) {
		link := r.URL.Query().Get("url")
		log.Println(link)
		short_link, err := logic.Get_short_link(link, storage)
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
