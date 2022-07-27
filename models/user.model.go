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

// gorm.Model 包括以下内容，这些都不用自己写，插入数据时会自动写入
// ID        uint `gorm:"primary_key"`
// CreatedAt time.Time
// UpdatedAt time.Time
// DeletedAt *time.Time `sql:"index"`
