package handler

import (
	"net/http"

	"github.com/0zyyy/money_record/helper"
	"github.com/0zyyy/money_record/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) FindAll(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user.ResponseFormatterUsers(users)})
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	users, err := h.userService.Login(input)
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "nerhasil login", "user": user.ResponseFormatterUser(users)})
}

func (h *userHandler) Register(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// cek apakah ada user yg sama
	already, err := h.userService.FindEmail(input)
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if already {
		response := helper.APIResponse("User already exist", http.StatusUnprocessableEntity, "failed", user.User{})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// register
	users, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Successfully created user", http.StatusOK, "success", user.ResponseFormatterUser(users))
	c.JSON(http.StatusOK, response)
}
