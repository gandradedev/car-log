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
	vehicleusecase "github.com/gaaandrade/car-log/internal/application/usecase/vehicle"
	"github.com/gaaandrade/car-log/internal/configuration/database"
	"github.com/gaaandrade/car-log/internal/configuration/swagger"
	"github.com/gaaandrade/car-log/internal/infrastructure/handler"
	infrarepo "github.com/gaaandrade/car-log/internal/infrastructure/repository"
	"github.com/gaaandrade/car-log/internal/infrastructure/routes"
)

func main() {
	db, err := database.New("./carlog.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	repo := infrarepo.NewVehicleRepository(db)

	createUC := vehicleusecase.NewCreateVehicleUseCase(repo)
	listUC := vehicleusecase.NewListVehiclesUseCase(repo)
	getUC := vehicleusecase.NewGetVehicleUseCase(repo)
	updateUC := vehicleusecase.NewUpdateVehicleUseCase(repo)
	deleteUC := vehicleusecase.NewDeleteVehicleUseCase(repo)
	updateKMUC := vehicleusecase.NewUpdateVehicleKMUseCase(repo)

	h := handler.NewVehicleHandler(createUC, listUC, getUC, updateUC, deleteUC, updateKMUC)

	mux := http.NewServeMux()
	routes.RegisterVehicleRoutes(mux, h)
	swagger.RegisterSwaggerRoutes(mux)

	log.Println("Server running at http://localhost:8080")
	log.Println("Swagger UI at http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
