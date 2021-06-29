package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UUID     string         `gorm:"primaryKey;not null;unique" json:"uuid"`
	Name     string         `gorm:"not null;unique" json:"name"`
	CreateAt *time.Time     `gorm:"autoCreateTime"  json:"-"`
	DeleteAt gorm.DeletedAt `json:"-"`
}

type UserRepository interface {
	CreateUser(*User) error
	GetUser(string) (*User, error)
}

type UserUsecase interface {
	CreateUser(*User) error
	GetUser(string) (*User, error)
}
