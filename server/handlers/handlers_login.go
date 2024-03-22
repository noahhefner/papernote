package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noahhefner/notes/middlewares"
	"fmt"
	//"noahhefner/notes/models"
)

// LoginHandler handles user login and issues JWT token
func Login(c *gin.Context) {

	/*
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	*/

	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Print(c.PostForm("username"))
	fmt.Print(c.PostForm("password"))

	if !middlewares.AuthenticateUser(username, password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := middlewares.GenerateJWT(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.Redirect(http.StatusFound, "/notes")

	//c.JSON(http.StatusOK, gin.H{"token": token})
}