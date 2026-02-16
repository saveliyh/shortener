package logic

import (
	"errors"
	"url_shortener/internal/database"
)

type TestStorage struct {
	added      bool
	added_link string
}

func (t *TestStorage) Connect_db() (database.Storage, error) {
	return t, nil
}

func (t *TestStorage) Check_long_link(long_link string) bool {
	return long_link == "long_link1" ||
		long_link == "long_link2" ||
		(t.added && long_link == "long_link3")
}

func (t *TestStorage) Check_short_link(short_link string) bool {
	return short_link == "short_link1" ||
		short_link == "short_link2" ||
		(t.added && short_link == t.added_link)
}

func (t *TestStorage) Get_short_link(long_link string) (string, error) {
	if long_link == "long_link1" {
		return "short_link1", nil
	}
	if long_link == "long_link2" {
		return "short_link2", nil
	}
	if long_link == "long_link3" && t.added {
		return t.added_link, nil
	}
	return "", errors.New("Long link not exist in database")
}

func (t *TestStorage) Get_long_link(short_link string) (string, error) {
	if short_link == "short_link1" {
		return "long_link1", nil
	}
	if short_link == "short_link2" {
		return "long_link2", nil
	}
	if short_link == t.added_link && t.added {
		return "long_link3", nil
	}
	return "", errors.New("Short link not exist in database")
}

func (t *TestStorage) Store_in_db(short_link string, long_link string) error {
	if short_link == "error" {
		return errors.New("Short link already exist in database")
	}
	t.added = true
	t.added_link = short_link
	return nil
}
