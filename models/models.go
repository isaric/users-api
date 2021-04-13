package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name string `json:"Name"`
}

type User struct {
	gorm.Model
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Group    Group  `gorm:"foreignKey:ID"`
}
