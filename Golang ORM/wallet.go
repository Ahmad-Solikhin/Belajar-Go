package Golang_ORM

import "time"

type Wallet struct {
	ID        string    `gorm:"primary_key"`
	UserId    string    `gorm:"column:user_id"`
	Balance   int64     `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
	User      *User     `gorm:"foreignkey:UserID;references:ID"`
}

func (w *Wallet) TableName() string {
	return "wallets"
}
