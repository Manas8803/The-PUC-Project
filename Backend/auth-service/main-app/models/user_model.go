package models

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type Register struct {
	OfficeName string `json:"office_name" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate:"required"`
}
