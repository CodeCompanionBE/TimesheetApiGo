package main

import (
	api "codecompanion/timesheet/api/internal"
	Env "codecompanion/timesheet/api/internal/env"
	OAuthHandlers "codecompanion/timesheet/api/internal/oauth/handlers"
	"fmt"
	"net/http"
)

func main() {
	// Start server and make db connection
	api := api.NewApi(http.NewServeMux())

	// Register routes
	OAuthHandlers.RegisterRoutes(api.Router)

	// Start server
	err := http.ListenAndServe(fmt.Sprintf(":%s", Env.Vars.Port), api.Router)

	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
	fmt.Println("Server listening on port 8080")
}
