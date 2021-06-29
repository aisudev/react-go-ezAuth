package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"react-go-auth/domain"
	authDelivery "react-go-auth/feature/auth/delivery"
	authRepo "react-go-auth/feature/auth/repository"
	authUsecase "react-go-auth/feature/auth/usecase"
	"react-go-auth/middleware"

	userDelivery "react-go-auth/feature/user/delivery"
	userRepo "react-go-auth/feature/user/repository"
	userUsecase "react-go-auth/feature/user/usecase"
)

var DB *gorm.DB

func init() {

	DB, _ = gorm.Open(sqlite.Open("db/db.db"), &gorm.Config{})

	AutoMigrate()

}

func main() {
	e := echo.New()

	auth := e.Group("/auth")
	authDelivery.NewAuthHandler(auth,
		authUsecase.NewAuthUsecase(
			authRepo.NewAuthRepository(DB),
		),
	)

	user := e.Group("/user")
	user.Use(middleware.AuthMiddleware(DB))
	userDelivery.NewUserHandler(user,
		userUsecase.NewUserUsecase(
			userRepo.NewUserRepository(DB),
		),
	)
	e.Start(":9999")
}

func AutoMigrate() {
	DB.AutoMigrate(&domain.Auth{}, &domain.User{})
}
