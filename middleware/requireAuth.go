package middleware

import (
	"BOOK-STORE/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main/initializers"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {

	//get the cookie from the request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return nil, nil
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []
		return []byte(os.Getenv("SECRET")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check if the token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(401, gin.H{"error": "Unauthorized"})
		}
		//check if the user exists
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(401)
		}

		//Attach the user to the context
		c.Set("user", user)
		//Continue with the request
		c.Next()
	} else {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
	}
}
