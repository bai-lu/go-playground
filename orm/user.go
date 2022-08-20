package orm

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Email   string
	IsAdmin bool
}
