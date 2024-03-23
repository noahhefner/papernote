package middlewares

import (
	"net/http"
	"time"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"noahhefner/notes/database"
)

var JWTSecret []byte

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

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
		if err != nil {
            c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
            return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})

		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		if !token.Valid {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func GenerateJWT(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}

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