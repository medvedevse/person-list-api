package entity

type NewPersonBody struct {
	Surname    string `json:"surname" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Patronymic string `json:"patronymic"`
}
