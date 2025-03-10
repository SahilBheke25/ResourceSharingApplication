package main

import (
	"log"
	"net/http"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/config"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/repository"

	_ "github.com/lib/pq"
)

func main() {

	sqlConfig := config.GetEnv()
	db := repository.InitializeDatabase(sqlConfig)
	defer db.Close()

	dependencies := app.InitializeDependencies(db)

	router := app.InitializeRoutes(dependencies)

	log.Fatal(http.ListenAndServe(":3000", router))
}
