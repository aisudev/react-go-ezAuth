package repository

import (
	"errors"
	"react-go-auth/domain"
	"react-go-auth/utils"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) CreateAuth(auth *domain.Auth) error {
	return r.db.Create(auth).Error
}

func (r *authRepository) GetAuth(filterBy string, filter string) (*domain.Auth, error) {

	query := "username = ?"
	if filterBy == "uuid" {
		query = "uuid = ?"
	}

	var auth domain.Auth
	if err := r.db.Where(query, filter).Find(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}

func (r *authRepository) UpdateAuth(uuid string, data map[string]interface{}) error {

	var auth domain.Auth
	if err := r.db.Model(&auth).Where("uuid = ?", uuid).Updates(data).Error; err != nil {
		utils.Log("", err)
		return err
	}

	return nil
}

func (r *authRepository) IsExist(filterBy, filter string) error {

	query := "username = ?"
	if filterBy == "uuid" {
		query = "uuid = ?"
	} else if filterBy == "refresh" {
		query = "refresh_token = ?"
	} else if filterBy == "access" {
		query = "access_token = ?"
	}

	var auth domain.Auth
	count := int64(0)

	if err := r.db.Model(&auth).Where(query, filter).Count(&count).Error; err != nil {
		return err
	}

	if count < 1 {
		return errors.New("not exist: " + filterBy)
	}

	return nil
}
