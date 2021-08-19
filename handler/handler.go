package handler

import (
	"fmt"
	"net/http"
	"startup/helper"
	"startup/users"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)



type userHandler struct{
	userService users.Service
}

func NewUserHandler(userService users.Service)*userHandler{
	return &userHandler{userService}
}

func (h *userHandler)RegisterUser(c *gin.Context){
	var input users.RegisterUserInput
	err:=c.ShouldBindJSON(&input)
	if err != nil{
		var errors []string
		for _, err := range err.(validator.ValidationErrors){
			errors=append(errors,err.Field()+":"+err.Tag())
		}
		errorMessage:=gin.H{"errors":errors}
		response :=helper.APIResponse("Register Account Failed",http.StatusUnprocessableEntity,"failed",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	newUser,err:=h.userService.RegisterUser(input)

	if err!=nil{
		response:=helper.APIResponse("Register Account Failed",http.StatusBadRequest,"failed",nil)
		 	c.JSON(http.StatusBadRequest,response)
			 return

	}
	
	token:="asasasass"
	formatter:=users.FormateUser(newUser,token)
	fmt.Println(formatter)
	response :=helper.APIResponse("account has been register",http.StatusOK,"succes",formatter)

	c.JSON(http.StatusOK,response)
}