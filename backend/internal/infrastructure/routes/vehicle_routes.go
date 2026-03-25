package routes

import (
	"net/http"

	"github.com/gaaandrade/car-log/internal/infrastructure/handler"
)

func RegisterVehicleRoutes(mux *http.ServeMux, h handler.VehicleHandler) {
	mux.HandleFunc("GET /vehicles", h.List)
	mux.HandleFunc("POST /vehicles", h.Create)
	mux.HandleFunc("GET /vehicles/{id}", h.GetByID)
	mux.HandleFunc("PUT /vehicles/{id}", h.Update)
	mux.HandleFunc("DELETE /vehicles/{id}", h.Delete)
	mux.HandleFunc("PATCH /vehicles/{id}/km", h.UpdateKM)
}
