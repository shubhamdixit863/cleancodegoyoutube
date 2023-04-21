package controllers

import (
	"cleancodego/internal/domain/services"
	"fmt"
)

type userController struct {
	userService services.UserService
}

func NewUserController(svc services.UserService) *userController {

	return &userController{
		svc,
	}
}

func (h *userController) SaveData() {

	fmt.Println(h.userService.SaveUser())
}
