// @title          Car Log API
// @version        1.0
// @description    REST API for managing vehicle maintenance history.

// @contact.name   Gabriel Andrade

// @host      localhost:8080
// @BasePath  /

package main

import (
	"log"
	"net/http"

	_ "github.com/gaaandrade/car-log/docs"
	"github.com/gaaandrade/car-log/internal/configuration/database"
	"github.com/gaaandrade/car-log/internal/configuration/swagger"
	"github.com/gaaandrade/car-log/internal/infrastructure/routes"
)

func main() {
	db, err := database.New("./carlog.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	routes.Register(mux)
	swagger.RegisterSwaggerRoutes(mux)

	log.Println("Server running at http://localhost:8080")
	log.Println("Swagger UI at http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
