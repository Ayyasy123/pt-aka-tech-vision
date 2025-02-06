package server

import (
	"github.com/Ayyasy123/pt-aka-tech-vision/internal/handler"
	"github.com/Ayyasy123/pt-aka-tech-vision/internal/repository"
	"github.com/Ayyasy123/pt-aka-tech-vision/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(db *gorm.DB, r *gin.Engine) {

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	UserHandler := handler.NewUserHandler(userUseCase)

	r.POST("/users", UserHandler.CreateUser)
}
