package routes

import (
	"database/sql"

	"github.com/andreybrumatti/crud-init/internal/controller"
	"github.com/andreybrumatti/crud-init/internal/repository"
	"github.com/andreybrumatti/crud-init/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	api := r.Group("/api/v1")

	setupUserRoutes(api, db)
}

func setupUserRoutes(r *gin.RouterGroup, db *sql.DB) {
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	ctrl := controller.NewUserController(svc)

	r.GET("/users", ctrl.GetAllUsers)
}