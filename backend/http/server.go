package http

import (
	"backend/http/middleware"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Server *http.Server
}

func (app *App) Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api_routes := http.NewServeMux()
	api_routes.Handle("/core/", http.StripPrefix("/core", CoreRoutes()))

	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", api_routes))

	middleware := middleware.CreateStack(
		middleware.Logging,
	)

	app.Server = &http.Server{
		Addr:    os.Getenv("HOST_ADDRESS") + ":" + os.Getenv("HOST_PORT"),
		Handler: middleware(router),
	}
}

func (app *App) Start() {
	log.Printf("Server listening at %s:%s", os.Getenv("HOST_ADDRESS"), os.Getenv("HOST_PORT"))
	app.Server.ListenAndServe()
}
