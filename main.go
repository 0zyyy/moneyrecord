package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/0zyyy/money_record/history"
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
	// userRepo := user.NewRepository(db)
	historyRepo := history.NewRepository(db)
	//init service
	// userService := user.NewService(userRepo)
	//nyoba
	ihir, err := historyRepo.FindAll()
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

	router.Run()
}
