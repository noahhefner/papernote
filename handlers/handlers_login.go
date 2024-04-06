package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noahhefner/notes/middlewares"
)

func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

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
	
}

func Logout(c *gin.Context) {

    // Clear the JWT token cookie
    c.SetCookie("jwt", "", -1, "/", "", false, true)

    // Redirect to login page
    c.Redirect(http.StatusTemporaryRedirect, "/login")
	
}

