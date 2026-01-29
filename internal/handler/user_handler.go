package handler

import (
	"net/http"
	"user-management-api/internal/dto"
	"user-management-api/internal/models"
	"user-management-api/internal/service"
	"user-management-api/internal/utils"
	"user-management-api/internal/validation"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

type GetUserByUUIDParam struct {
	UUID string `uri:"uuid" binding:"required,uuid"`
}

func (uh *UserHandler) GetAllUser(c *gin.Context) {
	users, err := uh.service.GetAllUser()
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	userDTOs := dto.MapUserToDTOList(users)
	utils.ResponseSuccess(c, http.StatusOK, userDTOs)
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidator(c, validation.HandleValidationErrors(err))
		
		return
	}

	user, errCreated := uh.service.CreateUser(user)
	if errCreated != nil {
		utils.ResponseError(c, errCreated)
		return		
	}

	userDTO := dto.MapUserToDTO(user)
	utils.ResponseSuccess(c, http.StatusCreated, &userDTO)
}

func (uh *UserHandler) GetUserByUUID(c *gin.Context) {
	var param GetUserByUUIDParam
	if err := c.ShouldBindUri(&param); err != nil {
		utils.ResponseValidator(c, validation.HandleValidationErrors(err))
		return
	}

	user, err := uh.service.GetUserByUUID(param.UUID)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	userDTO := dto.MapUserToDTO(user)
	utils.ResponseSuccess(c, http.StatusOK, &userDTO)
}
func (uh *UserHandler) UpdateUser(c *gin.Context) {}

func (uh *UserHandler) DeleteUser(c *gin.Context) {}