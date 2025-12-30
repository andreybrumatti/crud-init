package controller

import (
	"net/http"

	"github.com/andreybrumatti/crud-init/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc *service.UserService
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{svc: svc}
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.svc.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve users",
		})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No users found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"message": "Users retrieved successfully",
	})
}