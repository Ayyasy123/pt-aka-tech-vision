package handler

import (
	"net/http"

	"github.com/Ayyasy123/pt-aka-tech-vision/internal/model"
	"github.com/Ayyasy123/pt-aka-tech-vision/internal/response"
	"github.com/Ayyasy123/pt-aka-tech-vision/internal/usecase"
	"github.com/Ayyasy123/pt-aka-tech-vision/internal/validator"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	var req model.CreateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// perform validation
	if err := validator.ValidateStruct(&req); err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	user, err := u.userUseCase.CreateUser(&req)
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusCreated, "User created successfully", user)
}
