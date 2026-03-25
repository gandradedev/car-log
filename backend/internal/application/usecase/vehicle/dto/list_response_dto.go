package dto

// ListVehiclesResponseDTO is the output DTO for the full vehicle listing.
type ListVehiclesResponseDTO struct {
	Total    int                     `json:"total"`
	Vehicles []GetVehicleResponseDTO `json:"vehicles"`
}
