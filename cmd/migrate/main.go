package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		"postgres://your-username:your-pass@localhost:5432/db-name?sslmode=disable",
	)
	if err != nil {
		log.Fatal("error initialize migration")
	}

	cmd := os.Args[len(os.Args) - 1]

	if cmd == "up" {
		if err := m.Up(); err != nil && err == migrate.ErrNoChange {
			log.Println("migrate up error : ", err)
		} else {
			log.Println("migrate up success")
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err == migrate.ErrNoChange {
			log.Println("migrate down error : ", err)
		} else {
			log.Println("migrate down success")
		}
	}
}