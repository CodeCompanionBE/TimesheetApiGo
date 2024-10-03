package main

import (
	Env "codecompanion/timesheet/api/internal/env"
	OAuthHandlers "codecompanion/timesheet/api/internal/oauth/handlers"
	"fmt"
	"net/http"
)

func main() {
	// Start server and make db connection
	router := http.NewServeMux()

	// Register routes
	OAuthHandlers.RegisterRoutes(router)

	// Start server
	err := http.ListenAndServe(fmt.Sprintf(":%s", Env.Vars.Port), router)

	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
	fmt.Println("Server listening on port 8080")
}
