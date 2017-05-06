package model

import "github.com/jinzhu/gorm"

//User details
type User struct {
	gorm.Model
	UserName string `gorm:"not null;unique;index"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Phone    string
}
