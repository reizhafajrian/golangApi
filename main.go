package main

import (
	"log"
	"startup/handler"
	"startup/users"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main()  {
	dsn := "root:root@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		 log.Fatal(err.Error())
	}
	userRepsitory := users.NewRepository(db)
	userService := users.NewService(&userRepsitory)
	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()
	api:=router.Group("/api/v1")
	api.POST("/users",userHandler.RegisterUser)
	router.Run()



} 