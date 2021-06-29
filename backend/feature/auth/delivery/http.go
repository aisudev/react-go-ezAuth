package delivery

import (
	"fmt"
	"net/http"
	"react-go-auth/domain"
	"react-go-auth/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.AuthUsecase
}

func NewAuthHandler(e *echo.Group, usecase domain.AuthUsecase) *Handler {
	h := Handler{usecase: usecase}

	e.POST("", h.CreateAuthHandler)
	e.GET("", h.GetAuthHandler)

	e.POST("/refresh", h.RefreshAuthHandler)

	return &h
}

// CREATE AUTH HANDLER
func (h *Handler) CreateAuthHandler(c echo.Context) error {

	var auth domain.Auth

	if err := c.Bind(&auth); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, nil, nil, err))
	}

	if err := h.usecase.CreateAuth(&auth); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, nil, nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, nil, nil))
}

// GET AUTH HANDLER
func (h *Handler) GetAuthHandler(c echo.Context) error {

	var reqMap, resMap map[string]interface{}
	var err error

	if err = c.Bind(&reqMap); err != nil {
		utils.Log("", err)
		return c.JSON(http.StatusBadRequest, utils.Response(false, nil, nil, err))
	}

	fmt.Println(reqMap)

	if resMap, err = h.usecase.GetAuth(reqMap["username"].(string), reqMap["password"].(string)); err != nil {
		utils.Log("", err)
		return c.JSON(http.StatusBadRequest, utils.Response(false, nil, nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, resMap, nil))

}

// ACCESS AUTH HANDLER
// func (h *Handler) AccessAuthHandler(c echo.Context) error {

// 	reqMap := map[string]interface{}{}

// 	if err := c.Bind(&reqMap); err != nil {
// 		utils.Log("", err)
// 		return c.JSON(http.StatusBadRequest, utils.Response(false, nil, nil, err))
// 	}

// 	if err := h.usecase.VerifyAccessToken(reqMap["accessToken"].(string)); err != nil {
// 		utils.Log("", err)
// 		return c.JSON(http.StatusUnauthorized, utils.Response(false, nil, nil, err))
// 	}

// 	return c.JSON(http.StatusOK, utils.Response(true, nil, nil, nil))
// }

// REFRESH AUTH HANDLER
func (h *Handler) RefreshAuthHandler(c echo.Context) error {

	reqMap := map[string]interface{}{}
	resMap := map[string]interface{}{}
	var err error

	if err = c.Bind(&reqMap); err != nil {
		utils.Log("", err)
		return c.JSON(http.StatusBadRequest, utils.Response(false, nil, nil, err))
	}

	if resMap, err = h.usecase.VerifyRefreshToken(reqMap["refreshToken"].(string)); err != nil {
		utils.Log("", err)
		return c.JSON(http.StatusUnauthorized, utils.Response(false, nil, nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, resMap, nil))
}
