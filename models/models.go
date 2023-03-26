package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Id       uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Email    string `gorm:"primaryKey"`
	Name     string
	Students []*Student `gorm:"many2many:registers;"`
}

type Student struct {
	gorm.Model
	Id        uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Email     string `gorm:"primaryKey"`
	Name      string
	Suspended bool
}
