package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/0zyyy/money_record/auth"
	"github.com/0zyyy/money_record/handler"
	"github.com/0zyyy/money_record/helper"
	"github.com/0zyyy/money_record/history"
	"github.com/0zyyy/money_record/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Response struct {
	Msg string
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/money_record?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	// init repo
	userRepo := user.NewRepository(db)
	historyRepo := history.NewRepository(db)

	//init service
	userService := user.NewService(userRepo)
	historyService := history.NewService(*historyRepo)
	authService := auth.NewService()
	// init handler
	userHandler := handler.NewUserHandler(userService)
	historyHandler := handler.NewHistoryHandler(historyService)

	//api endpoint
	router := gin.Default()
	api := router.Group("api/v1")
	api.GET("/users", userHandler.FindAll)
	api.POST("/login", userHandler.Login)                                                               // login user
	api.POST("/register", userHandler.Register)                                                         // register user
	api.PUT("/history", authMiddleware(authService, userService), historyHandler.Update)                // update
	api.POST("/history", authMiddleware(authService, userService), historyHandler.Create)               // create history
	api.POST("/search/history", authMiddleware(authService, userService), historyHandler.SearchHistory) // search History
	api.POST("/search/income", authMiddleware(authService, userService), historyHandler.SearchIncome)   // search income
	api.POST("/anal", authMiddleware(authService, userService), historyHandler.Analysis)                // anal
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		var tokenString string
		token := strings.Split(authHeader, " ")
		if len(token) == 2 {
			tokenString = token[1]
		}
		validatedToken, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := validatedToken.Claims.(jwt.MapClaims)
		if !ok || !validatedToken.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))
		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
