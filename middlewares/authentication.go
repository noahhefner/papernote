package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"noahhefner/notes/database"
	"os"
	"time"
)

var JWTSecret []byte

/*
Attempt to reach JWT secret from environment variable. Use default JWT secret
if environment variable is not set.
*/
func InitJWTSecret() {

	JWTSecret = []byte(os.Getenv("JWT_SECRET"))

	if len(JWTSecret) == 0 {
		fmt.Println("JWT secret environment variable not set! Using default jwt secret.")
		JWTSecret = []byte("default_jwt_key")
	}

}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, err := c.Cookie("jwt")

		// cookie does not exist
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		token, err := validateJWT(tokenString)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		c.Set("username", claims["username"])
		c.Next()
	}
}

func GenerateJWT(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	fmt.Println(tokenString, err)

	return tokenString, nil
}

func AuthenticateUser(username string, password string) bool {

	user, err := database.GetUserByUsername(username)
	if err != nil {
		fmt.Print("user not found")
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Print("failed password compare")
		return false
	}

	return true
}

func validateJWT(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil

}
