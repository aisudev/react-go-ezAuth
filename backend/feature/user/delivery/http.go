package delivery

import (
	"net/http"
	"react-go-auth/domain"
	"react-go-auth/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(e *echo.Group, usecase domain.UserUsecase) *Handler {
	h := Handler{usecase: usecase}

	e.POST("", h.CreateUserHandler)
	e.GET("", h.GetUserHandler)

	return &h
}

func (h *Handler) CreateUserHandler(c echo.Context) error {

	var user domain.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, nil, nil, err))
	}

	user.UUID = c.Get("uuid").(string)

	if err := h.usecase.CreateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, nil, nil, err))
	}

	return c.JSON(http.StatusBadRequest, utils.Response(true, nil, nil, nil))

}

func (h *Handler) GetUserHandler(c echo.Context) error {

	reqMap := map[string]interface{}{}

	if err := c.Bind(&reqMap); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, nil, nil, err))
	}

	var user *domain.User
	var err error

	if user, err = h.usecase.GetUser(reqMap["uuid"].(string)); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, nil, nil, err))
	}

	return c.JSON(http.StatusBadRequest, utils.Response(true, nil, user, nil))

}
