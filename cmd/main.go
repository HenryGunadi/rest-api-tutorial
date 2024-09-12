package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/HenryGunadi/rest-api-tutorial/cmd/api"
	"github.com/HenryGunadi/rest-api-tutorial/config"
	"github.com/HenryGunadi/rest-api-tutorial/db"
)

func main() {
	// initialize db
	dbConn, err := db.NewPostgreSQL(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPORT,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName,
	))
	if err != nil {
		log.Fatal("error connecting to postgres")
	}

	if err := initDB(dbConn); err != nil {
		log.Fatal("connection with db error : ", err)
	}
	// start api server
	apiServer := api.NewAPIServer(":8080", dbConn)
	if err := apiServer.Run(); err != nil {
		log.Fatal("error running api server")
	}
}

func initDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to database : ", config.Envs.DBName)

	return nil
}