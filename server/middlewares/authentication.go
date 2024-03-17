package middlewares

import (
	"net/http"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"noahhefner/notes/database"
)

// JWTSecret is the secret key used to sign JWT tokens
var JWTSecret = []byte("secret_key") // Change this to a secure value

// Claims represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Middleware function to handle authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

// GenerateJWT generates a JWT token for the given username
func GenerateJWT(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}

	fmt.Print(tokenString)

	return tokenString, nil
}

// Check if username/password is valid
func AuthenticateUser(username string, password string) bool {

	user, err := database.GetUserByUsername(username)
	if err != nil {
		fmt.Print("user not found")
		return false
	}

	fmt.Print(user.Password)
	fmt.Print(password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Print("failed password compare")
		return false
	}

	return true
}