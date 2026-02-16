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

func initDB(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS links (short_link VARCHAR(10) PRIMARY KEY, long_link VARCHAR(100) UNIQUE)")
	return err
}

func (t Postgres) Connect_db() (Storage, error) {
	err := godotenv.Load("db.env")

	if err != nil {
		return t, err
	}
	log.Println("loaded .env")

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
	log.Println("Environment varibles extracted")

	host := "database"

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return t, err
	}
	log.Println("connected to DB")

	t.database = db
	err = initDB(db)
	if err != nil {
		return t, err
	}
	log.Println("DB initialized")
	return t, nil
}

func (t Postgres) Check_long_link(long_link string) bool {
	var count bool
	log.Println("check long link")
	err := t.database.QueryRow("SELECT EXISTS(SELECT * FROM links WHERE long_link = $1)", long_link).Scan(&count)
	if err != nil {

		return false
	}

	return count
}

func (t Postgres) Check_short_link(short_link string) bool {
	var count bool
	err := t.database.QueryRow("SELECT EXISTS(SELECT * FROM links WHERE short_link = $1)", short_link).Scan(&count)
	if err != nil {
		return false
	}
	return count
}

func (t Postgres) Get_short_link(long_link string) (string, error) {
	var short_link string
	err := t.database.QueryRow("SELECT short_link FROM links WHERE long_link = $1", long_link).Scan(&short_link)
	if err != nil {
		return "", err
	}
	return short_link, nil
}

func (t Postgres) Get_long_link(short_link string) (string, error) {
	var long_link string
	err := t.database.QueryRow("SELECT long_link FROM links WHERE short_link = $1", short_link).Scan(&long_link)
	if err != nil {
		return "", err
	}
	return long_link, nil
}

func (t Postgres) Store_in_db(short_link string, long_link string) error {
	_, err := t.database.Exec("INSERT INTO links (short_link, long_link) VALUES ($1, $2)", short_link, long_link)

	return err
}

func (t Postgres) Close() {
	t.database.Close()
}
