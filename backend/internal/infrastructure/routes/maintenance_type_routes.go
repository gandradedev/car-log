package routes

import (
	"net/http"

	"github.com/gaaandrade/car-log/internal/infrastructure/handler"
)

func RegisterMaintenanceTypeRoutes(mux *http.ServeMux, h handler.MaintenanceTypeHandler) {
	mux.HandleFunc("GET /maintenance-types", h.List)
	mux.HandleFunc("POST /maintenance-types", h.Create)
	mux.HandleFunc("PUT /maintenance-types/{id}", h.Update)
	mux.HandleFunc("DELETE /maintenance-types/{id}", h.Delete)
}
