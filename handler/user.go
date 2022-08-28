package handler

import (
	"net/http"

	"github.com/0zyyy/money_record/auth"
	"github.com/0zyyy/money_record/helper"
	"github.com/0zyyy/money_record/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) FindAll(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user.ResponseFormatterUsers(users, "1111")})
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Failed to login", http.StatusUnprocessableEntity, "failed", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// token generate dulu
	users, err := h.userService.Login(input)
	if err != nil {
		response := helper.APIResponse("Failed to login", http.StatusUnprocessableEntity, "failed", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(users.IDUser)
	if err != nil {
		response := helper.APIResponse("Failed to login", http.StatusUnprocessableEntity, "failed", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Berhasil login", http.StatusOK, "success", user.ResponseFormatterUser(users, token))
	c.JSON(http.StatusOK, response)
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
	token, err := h.authService.GenerateToken(users.IDUser)
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Successfully created user", http.StatusOK, "success", user.ResponseFormatterUser(users, token))
	c.JSON(http.StatusOK, response)
}
