package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"noahhefner/notes/middlewares"
)

type loginResponse struct {
	Error bool
}

/*
Validate user credentials, set jwt on successful login.
*/
func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if !middlewares.AuthenticateUser(username, password) {

		context := loginResponse{
			Error: true,
		}

		c.HTML(http.StatusUnauthorized, "login.html", context)
		return
	}

	token, err := middlewares.GenerateJWT(username)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.Redirect(http.StatusSeeOther, "/notes")

}

/*
Clear JWT in the clients browser and redirect to login page.
*/
func Logout(c *gin.Context) {

	// Clear the JWT token cookie
	c.SetCookie("jwt", "", -1, "/", "", false, true)

	// Client-side redirect to login page
	c.Header("HX-Redirect", "/login")

}
