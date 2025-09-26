package middleware

import (
	"fmt"
	"go-auth/database"
	"strings"

	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(db *gorm.DB) {
	DB = db
}
func RequireAuth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/login") {
			c.Next()
			return
		}
		token2 := c.Request.Header.Get("token")
		tokenString, err := c.Cookie("token")

		if tokenString == "" || err != nil {
			if token2 != "" {
				tokenString = token2
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthenticated")
				return
			}
		}
		//Within the parsing function, it verifies that the token is signed with the expected signing method 
		//(HMAC) and retrieves the secret key from the environment variable named "SECRET".
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})
		//If the token is valid, the middleware checks if the token has expired by 
		//comparing the current time with the expiration time (exp) from the token's claims.
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			email := claims["email"].(string) //It extracts the user's email address from the token's claims.
			

			var user database.User
			result := DB.First(&user, "email = ?", email)
			if result.RowsAffected == 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized") //If the user's email is empty, the middleware responds with a 401 Unauthorized status.
				c.Abort()
				return
			} else {
				if user.Role != role {
					c.AbortWithStatusJSON(http.StatusForbidden, "Role mismatch you are not "+role+" of library ")
					c.Abort()
					return
				}
			}
			if user.Email == "" {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}


			// Attach the request
			c.Set("user", user)

			c.Next()  //to proceed to the next middleware or route handler
		}
	}
}
