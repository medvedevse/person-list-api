package entity

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Surname     string `json:"surname"`
	Name        string `json:"name"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
	Gender      string `json:"gender"`
}

type PersonAge struct {
	Age int `json:"age"`
}

type PersonGender struct {
	Gender string `json:"gender"`
}
