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
	"os"
	"strconv"

	"github.com/joho/godotenv"

	_ "github.com/gaaandrade/car-log/docs"
	maintenancetypeusecase "github.com/gaaandrade/car-log/internal/application/usecase/maintenance_type"
	vehicleusecase "github.com/gaaandrade/car-log/internal/application/usecase/vehicle"
	"github.com/gaaandrade/car-log/internal/configuration/database"
	"github.com/gaaandrade/car-log/internal/configuration/swagger"
	"github.com/gaaandrade/car-log/internal/infrastructure/handler"
	infrarepo "github.com/gaaandrade/car-log/internal/infrastructure/repository"
	"github.com/gaaandrade/car-log/internal/infrastructure/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./carlog.db"
	}

	runStartupScripts, err := strconv.ParseBool(os.Getenv("EXECUTE_STARTUP_SCRIPTS"))
	if err != nil {
		log.Printf("invalid EXECUTE_STARTUP_SCRIPTS value, defaulting to false: %v", err)
		runStartupScripts = false
	}

	db, err := database.New(database.Config{
		Path:              dbPath,
		RunStartupScripts: runStartupScripts,
	})
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

	mtRepo := infrarepo.NewMaintenanceTypeRepository(db)

	mtCreateUC := maintenancetypeusecase.NewCreateMaintenanceTypeUseCase(mtRepo)
	mtListUC := maintenancetypeusecase.NewListMaintenanceTypesUseCase(mtRepo)
	mtUpdateUC := maintenancetypeusecase.NewUpdateMaintenanceTypeUseCase(mtRepo)
	mtDeleteUC := maintenancetypeusecase.NewDeleteMaintenanceTypeUseCase(mtRepo)

	mtHandler := handler.NewMaintenanceTypeHandler(mtCreateUC, mtListUC, mtUpdateUC, mtDeleteUC)

	mux := http.NewServeMux()
	routes.RegisterVehicleRoutes(mux, h)
	routes.RegisterMaintenanceTypeRoutes(mux, mtHandler)
	swagger.RegisterSwaggerRoutes(mux)

	log.Println("Server running at http://localhost:8080")
	log.Println("Swagger UI at http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
