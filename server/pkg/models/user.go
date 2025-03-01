package models

import "time"

type User struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UserName  string    `gorm:"column:user_name" json:"user_name"`
}
