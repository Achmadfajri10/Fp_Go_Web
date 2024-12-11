package middlewares

import (
	"Fp_Go_Web/config"
	"Fp_Go_Web/entities"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAuth(c *gin.Context) {

	tokenString, err := c.Cookie("jwt")
	if err != nil {
		if err == http.ErrNoCookie {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user entities.User
	config.DB.Where("ID=?", claims["id"]).Find(&user)

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.SetCookie("currentUser", user.Username, 3600, "/", "", true, true)

	c.Next()

}