package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikhailvzhzhv/go-restorio/user-service/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler *UserHandler) UserLogin(ctx *gin.Context) {
	// ctx.ShouldBindJSON()
}
