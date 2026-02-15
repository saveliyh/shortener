package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func (t Postgres) Connect_db() (Storage, error) {
	err := godotenv.Load("db.env")

	if err != nil {
		return t, err
	}
	log.Println("loaded .env")

	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		return t, errors.New("POSTGRES_HOST not set")
	}
	port, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		return t, errors.New("POSTGRES_PORT not set")
	}
	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		return t, errors.New("POSTGRES_USER not set")
	}
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		return t, errors.New("POSTGRES_PASSWORD not set")
	}
	dbname, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		return t, errors.New("POSTGRES_DB not set")
	}

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return t, err
	}
	t.Database = db
	return t, nil
}

func (t Postgres) Check_long_link(long_link string) bool {
	var count int
	err := t.Database.QueryRow("SELECT EXISTS(SELECT * FROM shortener WHERE long_link = $1)", long_link).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (t Postgres) Check_short_link(short_link string) bool {
	var count int
	err := t.Database.QueryRow("SELECT EXISTS(SELECT * FROM shortener WHERE short_link = $1)", short_link).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (t Postgres) Get_short_link(long_link string) (string, error) {
	var short_link string
	err := t.Database.QueryRow("SELECT short_link FROM shortener WHERE long_link = $1", long_link).Scan(&short_link)
	if err != nil {
		return "", err
	}
	return short_link, nil
}

func (t Postgres) Get_long_link(short_link string) (string, error) {
	var long_link string
	err := t.Database.QueryRow("SELECT long_link FROM shortener WHERE short_link = $1", short_link).Scan(&long_link)
	if err != nil {
		return "", err
	}
	return long_link, nil
}

func (t Postgres) Store_in_db(short_link string, long_link string) error {
	_, err := t.Database.Exec("INSERT INTO shortener (short_link, long_link) VALUES ($1, $2)", short_link, long_link)
	if err != nil {
		return err
	}
	return nil
}
