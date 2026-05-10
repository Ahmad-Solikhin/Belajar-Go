package Golang_ORM

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserId      string `gorm:"column:user_id;not null"`
	Title       string `gorm:"column:title;not null"`
	Description string `gorm:"column:description;not null"`
}

func (t *Todo) TableName() string {
	return "todos"
}
