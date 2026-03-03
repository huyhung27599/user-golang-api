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

type GetUsersParams struct {
	Search string `form:"search" binding:"omitempty,min=3,max=100,search"`
	Page int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}
	
func (uh *UserHandler) GetAllUser(c *gin.Context) {
	var params GetUsersParams
	if err := c.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidator(c, validation.HandleValidationErrors(err))
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	users, err := uh.service.GetAllUser(params.Search, params.Page, params.Limit)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	userDTOs := dto.MapUserToDTOList(users)
	utils.ResponseSuccess(c, http.StatusOK, userDTOs)
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var input dto.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(c, validation.HandleValidationErrors(err))
		
		return
	}

	user := input.MapCreateUserInputToUserModel()

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
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var param GetUserByUUIDParam
	if err := c.ShouldBindUri(&param); err != nil {
		utils.ResponseValidator(c, validation.HandleValidationErrors(err))
		return
	}
	var input dto.UpdateUserInput


	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(c, validation.HandleValidationErrors(err))
		
		return
	}
	user := input.MapUpdateUserInputToUserModel()

	user, errUpdated := uh.service.UpdateUser(param.UUID, user)
	if errUpdated != nil {
		utils.ResponseError(c, errUpdated)
		return
	}

	userDTO := dto.MapUserToDTO(user)
	utils.ResponseSuccess(c, http.StatusOK, &userDTO)
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	var param GetUserByUUIDParam
	if err := c.ShouldBindUri(&param); err != nil {
		utils.ResponseValidator(c, validation.HandleValidationErrors(err))
		return
	}

	if errDeleted := uh.service.DeleteUser(param.UUID); errDeleted != nil {
		utils.ResponseError(c, errDeleted)
		return
	}

	utils.ResponseSuccess(c, http.StatusNoContent, nil)
}