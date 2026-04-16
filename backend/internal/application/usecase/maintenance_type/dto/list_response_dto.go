package dto

// ListMaintenanceTypesResponseDTO is the output DTO for the full maintenance type listing.
type ListMaintenanceTypesResponseDTO struct {
	Total            int                             `json:"total"`
	MaintenanceTypes []GetMaintenanceTypeResponseDTO `json:"maintenance_types"`
}
