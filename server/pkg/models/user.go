package models

import (
	"database/sql"
)

type User struct {
	ID        int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt sql.NullTime `gorm:"column:created_at" json:"created_at"`
	UserName  string       `gorm:"column:user_name" json:"user_name"`
}
