package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Account  string `json:"account"`
	Username string `json:"username"`
	Password string `json:"password"`
}
