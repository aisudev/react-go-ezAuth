package middleware

import (
	"errors"
	"net/http"
	"react-go-auth/domain"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var jwtKey = []byte("secret")

type Claim struct {
	Username string
	Secret   string
	ExpireAt int64
	jwt.StandardClaims
}

func AuthMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token := c.Request().Header.Get("Authorization")

			tokenSplit := strings.Split(token, " ")

			var err error
			var uuid string

			if uuid, err = VerifyAccessToken(db, tokenSplit[1]); err != nil {
				return c.String(http.StatusUnauthorized, err.Error())
			}

			c.Set("uuid", uuid)
			return next(c)
		}
	}
}

// *Access
func VerifyAccessToken(db *gorm.DB, token string) (string, error) {

	claim, err := VerifyJWToken(token)

	if err != nil {
		return "", err
	}

	if err := IsAuthExist(db, "access", token); err != nil {
		return "", err
	}

	if err := IsAuthExist(db, "username", claim.Username); err != nil {
		return "", err
	}

	return claim.Secret, nil

}

// *JWT
func GenerateJWToken(username, secret string, minute int64) (string, error) {

	token := jwt.New(jwt.SigningMethodHS384)

	claim := token.Claims.(jwt.MapClaims)
	claim["Username"] = username
	claim["Secret"] = secret
	claim["ExpireAt"] = time.Now().Add(time.Minute * time.Duration(minute)).Unix()

	return token.SignedString(jwtKey)

}

func VerifyJWToken(token string) (*Claim, error) {

	claim := &Claim{}

	tkn, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		return nil, errors.New("token's not valid.")
	}

	if err != nil {
		return nil, err
	}

	if claim.ExpireAt-time.Now().Unix() <= 0 {
		return nil, errors.New("token's expired.")
	}

	return claim, nil

}

func IsAuthExist(db *gorm.DB, filterBy, filter string) error {
	var auth domain.Auth
	count := int64(0)

	query := "username = ?"
	if filterBy == "uuid" {
		query = "uuid = ?"
	} else if filterBy == "refresh" {
		query = "refresh_token = ?"
	} else if filterBy == "access" {
		query = "access_token = ?"
	}

	if err := db.Model(&auth).Where(query, filter).Count(&count).Error; err != nil {
		return err
	}

	if count < 1 {
		return errors.New("not exist: " + filterBy)
	}

	return nil
}
