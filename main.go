package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/0zyyy/money_record/handler"
	"github.com/0zyyy/money_record/history"
	"github.com/0zyyy/money_record/user"
	"github.com/gin-gonic/gin"
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
	// init handler
	userHandler := handler.NewUserHandler(userService)
	historyHandler := handler.NewHistoryHandler(historyService)
	//nyoba
	ihir, err := historyService.Analysis("2022-06-20")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ihir)
	}
	router := gin.Default()
	api := router.Group("api/v1")
	api.GET("/hello", func(c *gin.Context) {
		response := Response{Msg: "Hemlo"}
		c.JSON(http.StatusOK, response)
	})
	api.GET("/users", userHandler.FindAll)
	api.POST("/login", userHandler.Login)
	api.POST("/register", userHandler.Register)
	api.PUT("/history", historyHandler.Update)  // update
	api.POST("/history", historyHandler.Create) // create history

	router.Run()
}
