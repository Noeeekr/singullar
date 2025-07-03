package handlers

type SelectByInstitutionRequest struct {
	InstitutionId int `json:"institution_id" binding:"required"`
}
