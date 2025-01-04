package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Contrasena string `json:"contrasena"`
	Email      string `json:"email"`
}

func (User) TableName() string {
	return "User"
}