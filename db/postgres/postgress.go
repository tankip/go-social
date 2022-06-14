package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("postgres", "postgres://postgres:postgrespw@localhost:55004/go-social?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = Db.Ping(); err != nil {
		panic(err)
	}
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := postgres.WithInstance(Db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
