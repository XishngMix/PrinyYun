package Modules

import (
	_ "github.com/jinzhu/gorm"
)

type Province struct {
	Id int `gorm:"AUTO_INCREMENT;primary_key"`
	Name string `gorm:"type:varchar(32);not null;unique"`
}

type City struct {
	Id int `gorm:"AUTO_INCREMENT;primary_key"`
	Name string `gorm:"type:varchar(32);not null;unique"`
}

type Area struct {
	Id int `gorm:"AUTO_INCREMENT;primary_key"`
	Name string `gorm:"type:varchar(128);not null;unique"`
}