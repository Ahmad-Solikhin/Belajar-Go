package Golang_ORM

import "time"

type GuestBook struct {
	ID        int64     `gorm:"primary_key;auto_increment"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null"`
	Message   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime;autoUpdateTime"`
}

func (g *GuestBook) TableName() string {
	return "guest_books"
}
