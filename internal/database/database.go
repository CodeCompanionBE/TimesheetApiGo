package Db

import (
	Env "codecompanion/timesheet/api/internal/env"
	"database/sql"
	"fmt"
)

func initDb() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Env.Vars.DbHost, Env.Vars.DbPort, Env.Vars.DbUser, Env.Vars.DbPassword, Env.Vars.DbName))

	if err != nil {
		fmt.Println("Error opening database connection")
		panic(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Error pinging database connection")
		panic(pingErr)
	}

	return db
}

var Db *sql.DB = initDb()
