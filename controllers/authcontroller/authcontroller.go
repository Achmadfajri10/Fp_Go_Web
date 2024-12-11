package authcontroller

import (
	"Fp_Go_Web/entities"
	"Fp_Go_Web/models/authmodel"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Add(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "register.html", nil)
		return
	}

	if c.Request.Method == "POST" {
		var authInput entities.RegisterInput
		authInput.Username = c.PostForm("username")
		authInput.Email = c.PostForm("email")
		authInput.Password = c.PostForm("password")

		userFound, err := authmodel.FindUserByUsername(authInput.Username)

		if userFound.ID != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already used"})
			return
		}

		emailFound, err := authmodel.FindUserByUsername(authInput.Email)

		if emailFound.ID != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already used"})
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := entities.User{
			Username: authInput.Username,
			Email:    authInput.Email,
			Password: string(passwordHash),
		}

		if ok := authmodel.Create(user); !ok {
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Failed to create category"})
			return
		}

		c.Redirect(http.StatusSeeOther, "/login/")
	}
}

func GetUserProfile(c *gin.Context) {
	user, err := c.Cookie("currentUser")
	if err == nil {
        c.JSON(http.StatusOK, gin.H{"user": user})
    } else {
        c.JSON(http.StatusOK, gin.H{"user": nil})
    }
}

func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}

	if c.Request.Method == "POST" {
		var loginInput entities.LoginInput
		loginInput.Identifier = c.PostForm("loginInput")
		loginInput.Password = c.PostForm("password")

		userFound, err := authmodel.FindUserByUsername(loginInput.Identifier)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong username/email or password"})
			return
		}
		if userFound.ID == 0 {
			emailFound, err := authmodel.FindUserByUsername(loginInput.Identifier)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong username/email or password"})
				return
			}
			userFound = emailFound
		}

		if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginInput.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
			return
		}

		generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  userFound.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
		}

		c.SetCookie("jwt", token, 3600, "/", "", true, true)

		c.Redirect(http.StatusSeeOther, "/products")
	}
}

func Logout(c *gin.Context) {
    c.SetCookie("jwt", "", -1, "/", "", true, true)
	c.SetCookie("currentUser", "", -1, "/", "", true, true)

    c.Redirect(http.StatusSeeOther, "/login")
}