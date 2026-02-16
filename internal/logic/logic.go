package logic

import (
	"math/rand"
	"url_shortener/internal/database"
)

func Unshorten_link(short_link string, storage database.Storage) (string, error) {

	long_link, err := storage.Get_long_link(short_link)
	if err != nil {
		return "", err
	}
	return long_link, nil
}

func Get_short_link(long_link string, storage database.Storage) (string, error) {
	if storage.Check_long_link(long_link) {
		return storage.Get_short_link(long_link)
	}
	var short_link string
	for {

		short_link = create_short_link()

		if !storage.Check_short_link(short_link) {
			break
		}
	}
	err := storage.Store_in_db(short_link, long_link)
	if err != nil {
		if storage.Check_long_link(long_link) {
			return storage.Get_short_link(long_link)
		}
		return "", err
	}
	return short_link, nil
}

func create_short_link() string {
	short_link := make([]byte, LINK_LENGTH)
	for i := range LINK_LENGTH {
		short_link[i] = ALLOWED_CHARS[rand.Intn(len(ALLOWED_CHARS))]
	}

	return string(short_link)
}
