package usecase

import (
	"react-go-auth/domain"
	"react-go-auth/middleware"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authUsercase struct {
	repo domain.AuthRepository
}

func NewAuthUsecase(repo domain.AuthRepository) domain.AuthUsecase {
	return authUsercase{
		repo: repo,
	}
}

func (u authUsercase) CreateAuth(auth *domain.Auth) error {

	auth.UUID = GenerateUUID()
	auth.Password = GeneratePasswordHash(auth.Password)

	return u.repo.CreateAuth(auth)
}

func (u authUsercase) GetAuth(username, password string) (map[string]interface{}, error) {

	var auth *domain.Auth
	var err error

	if auth, err = u.repo.GetAuth("username", username); err != nil {
		return nil, err
	}

	if err := ComparePasswordHash(password, auth.Password); err != nil {
		return nil, err
	}

	tokenMap := map[string]interface{}{}
	var accessToken string
	var refreshToken string

	if accessToken, err = middleware.GenerateJWToken(auth.Username, auth.UUID, 1); err != nil {
		return nil, err
	}
	tokenMap["access_token"] = accessToken

	if refreshToken, err = middleware.GenerateJWToken(auth.Username, auth.Password, 5); err != nil {
		return nil, err
	}
	tokenMap["refresh_token"] = refreshToken

	if err = u.repo.UpdateAuth(auth.UUID, tokenMap); err != nil {
		return nil, err
	}

	return tokenMap, nil
}

func (u authUsercase) VerifyAccessToken(token string) error {

	claim, err := middleware.VerifyJWToken(token)

	if err != nil {
		return err
	}

	if err := u.repo.IsExist("access", token); err != nil {
		return err
	}

	if err := u.repo.IsExist("username", claim.Username); err != nil {
		return err
	}

	return nil
}

func (u authUsercase) VerifyRefreshToken(token string) (map[string]interface{}, error) {

	claim, err := middleware.VerifyJWToken(token)

	if err != nil {
		return nil, err
	}

	if err := u.repo.IsExist("refresh", token); err != nil {
		return nil, err
	}

	if err := u.repo.IsExist("uuid", claim.Secret); err != nil {
		return nil, err
	}

	tokenMap := map[string]interface{}{}
	var accessToken string
	var refreshToken string
	var auth *domain.Auth

	if auth, err = u.repo.GetAuth("uuid", claim.Secret); err != nil {
		return nil, err
	}

	if accessToken, err = middleware.GenerateJWToken(auth.Username, auth.UUID, 1); err != nil {
		return nil, err
	}
	tokenMap["access_token"] = accessToken

	if refreshToken, err = middleware.GenerateJWToken(auth.Username, auth.Password, 5); err != nil {
		return nil, err
	}
	tokenMap["refresh_token"] = refreshToken

	if err = u.repo.UpdateAuth(auth.UUID, tokenMap); err != nil {
		return nil, err
	}

	return tokenMap, nil
}

// *HASH
func GeneratePasswordHash(password string) string {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(hashPassword)
}

func ComparePasswordHash(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

// *UUID
func GenerateUUID() string {
	_uuid, _ := uuid.DefaultGenerator.NewV4()
	return _uuid.String()
}
