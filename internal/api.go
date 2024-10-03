package api

import (
	Env "codecompanion/timesheet/api/internal/env"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type API struct {
	Router *http.ServeMux
	DB     *sql.DB
}

func NewApi(router *http.ServeMux) *API {
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

	return &API{
		Router: router,
		DB:     db,
	}
}
