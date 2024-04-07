package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"noahhefner/notes/database"
	"noahhefner/notes/models"
)

func AddUser(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	// Check if the username already exists
	if database.UserExists(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	// Hash the users password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	password = string(hashedPassword)

	newUser := models.User{
		Username: username,
		Password: password,
	}

	// Insert the user into the database
	if err := database.InsertUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to the /notes page
	c.Redirect(http.StatusFound, "/login")

}

func GetUser(c *gin.Context) {

	username := c.Param("username")

	// Retrieve the user from the database
	user, err := database.GetUserByUsername(username)
	if err != nil {
		if err == database.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the user details
	c.JSON(http.StatusOK, user)

}
