package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Id       uint64 `gorm:"primaryKey"`
	Email    string `gorm:"primaryKey"`
	Name     string
	Students []*Student `gorm:"many2many:registers;"`
}

type Student struct {
	gorm.Model
	Id        uint64 `gorm:"primaryKey"`
	Email     string `gorm:"primaryKey"`
	Name      string
	Suspended bool
}
