package controllers

import (
	"clean/domain"
	"clean/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUsecase domain.UserUsecase
}

func NewAuthController(a domain.UserUsecase) *AuthController {
	return &AuthController{authUsecase: a}
}

func (c *AuthController) Register(ctx echo.Context) error {
	user := new(domain.User)
	if err := ctx.Bind(user); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "Gagal menginput data: "+err.Error())
	}

	if user.Email == "" || user.Password == "" {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "Email dan password harus diisi")
	}

	if err := c.authUsecase.Register(user); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal mendaftar pengguna: "+err.Error())
	}

	return helper.JSONSuccessResponse(ctx, "Berhasil register")
}

func (c *AuthController) Login(ctx echo.Context) error {
	user := new(domain.User)
	if err := ctx.Bind(user); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "Gagal menginput data: "+err.Error())
	}

	if user.Email == "" || user.Password == "" {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "Email dan password harus diisi")
	}

	token, err := c.authUsecase.Login(user.Email, user.Password)
	if err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusUnauthorized, "Gagal mendapatkan token: "+err.Error())
	}

	return helper.JSONSuccessResponse(ctx, map[string]string{"token": token})
}
