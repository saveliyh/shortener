package main

import (
	"log"
	"math/rand"
	"url_shortener/database"
)

const ALLOWED_CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const LINK_LENGTH = 10

func unshorten_link(shorten_link string, storage database.Storage) (string, error) {

	long_link, err := storage.Get_long_link(shorten_link)
	if err != nil {
		return "", err
	}
	return long_link, nil
}

func get_short_link(long_link string, storage database.Storage) (string, error) {
	log.Println("in get_short_link")
	if storage.Check_long_link(long_link) {
		return storage.Get_short_link(long_link)
	}
	var short_link string
	for {

		short_link = create_short_link(long_link)

		if !storage.Check_short_link(short_link) {
			break
		}
	}
	var err = storage.Store_in_db(short_link, long_link)
	if err != nil {
		return "", err
	}
	return short_link, nil
}

// TODO: redo to hash
func create_short_link(long_link string) string {
	short_link := make([]byte, LINK_LENGTH)
	for i := range LINK_LENGTH {
		short_link[i] = ALLOWED_CHARS[rand.Intn(len(ALLOWED_CHARS))]
	}

	return string(short_link)
}
