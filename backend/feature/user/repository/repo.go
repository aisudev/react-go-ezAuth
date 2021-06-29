package repository

import (
	"react-go-auth/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *domain.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetUser(uuid string) (*domain.User, error) {

	var user *domain.User

	if err := r.db.Where("uuid = ?", uuid).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
