package usecase

import "react-go-auth/domain"

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserUsecase) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) CreateUser(user *domain.User) error {
	return u.repo.CreateUser(user)
}

func (u *userUsecase) GetUser(uuid string) (*domain.User, error) {
	return u.repo.GetUser(uuid)
}
