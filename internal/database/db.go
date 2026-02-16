package database

import (
	"database/sql"
	"sync"
)

type Storage interface {
	Connect_db() (Storage, error)
	Get_long_link(short_link string) (string, error)
	Get_short_link(long_link string) (string, error)
	Store_in_db(short_link string, long_link string) error
	Check_long_link(long_link string) bool
	Check_short_link(short_link string) bool
}

type Postgres struct {
	database *sql.DB
}
type InMemory struct {
	longToShort map[string]string
	shortToLong map[string]string
	sync.RWMutex
}
