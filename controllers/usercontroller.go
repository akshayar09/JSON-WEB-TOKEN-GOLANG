package controllers

import (
	"net/http"
	"restAPI/database"
	"restAPI/model"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	// Whatever is sent by the client as a JSON body will be mapped into the user variable
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// Hash the password using the bcrypt helpers
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// Once hashed, we store the user data into the database using the GORM
	record := database.DB.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	// If everything goes well, we send back  to the client
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.UserName})
}
