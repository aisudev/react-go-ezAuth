package domain

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	UUID         string         `gorm:"primaryKey;not null;unique" json:"uuid"`
	Username     string         `gorm:"not null;unique" json:"username"`
	Password     string         `gorm:"not null" json:"password"`
	AccessToken  string         `gorm:"null" json:"accessToken"`
	RefreshToken string         `gorm:"null" json:"refreshToken"`
	CreateAt     *time.Time     `gorm:"autoCreateTime"  json:"-"`
	DeleteAt     gorm.DeletedAt `json:"-"`
}

type AuthRepository interface {
	CreateAuth(*Auth) error
	GetAuth(string, string) (*Auth, error)
	UpdateAuth(string, map[string]interface{}) error
	IsExist(string, string) error
}

type AuthUsecase interface {
	CreateAuth(*Auth) error
	GetAuth(string, string) (map[string]interface{}, error)
	VerifyAccessToken(string) error
	VerifyRefreshToken(string) (map[string]interface{}, error)
}
