package main

import "fmt"

var shortToLong = make(map[string]string)
var longToShort = make(map[string]string)

func unshorten_link(shorten_link string) string {
	long_link := get_from_db(shorten_link)
	return long_link
}

func get_short_link(long_link string) string {
	short_link := longToShort[long_link]
	if short_link != "" {
		return longToShort[long_link]
	}
	short_link = create_short_link(long_link)
	store_in_db(short_link, long_link)
	return short_link
}

var counter = 0

func create_short_link(long_link string) string {
	short_link := "link" + fmt.Sprint(counter)
	counter++
	return short_link
}

func get_from_db(short_link string) string {
	return shortToLong[short_link]
}

func store_in_db(short_link string, long_link string) {
	shortToLong[short_link] = long_link
	longToShort[long_link] = short_link
}
