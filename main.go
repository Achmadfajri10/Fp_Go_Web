package main

import (
	"Fp_Go_Web/config"
	homecontroller "Fp_Go_Web/controllers"
	"Fp_Go_Web/controllers/authcontroller"
	"Fp_Go_Web/controllers/categorycontroller"
	"Fp_Go_Web/controllers/productcontroller"
	"Fp_Go_Web/middlewares"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization"}
	r.Use(cors.New(config))

	r.Static("/static", "./static")

	r.LoadHTMLGlob("views/**/*.html")

	//1.Homepage
	r.GET("/", homecontroller.Welcome)
	r.GET("/register", authcontroller.Add)
	r.POST("/register", authcontroller.Add)
	r.GET("/login", authcontroller.Login)
	r.POST("/login", authcontroller.Login)
	r.GET("/logout", authcontroller.Logout)
	r.GET("/user", authcontroller.GetUserProfile)
	r.GET("/profile", authcontroller.EditProfile)
	r.POST("/profile", authcontroller.EditProfile)
	r.GET("/delete", authcontroller.Delete)

	//2.Categories
	categories := r.Group("/categories", middlewares.CheckAuth)
	{
		categories.GET("/", categorycontroller.Index)
		categories.GET("/add", categorycontroller.Add)
		categories.POST("/add", categorycontroller.Add)
		categories.GET("/edit", categorycontroller.Edit)
		categories.POST("/edit", categorycontroller.Edit)
		categories.GET("/delete", categorycontroller.Delete)
	}

	//3.Products
	products := r.Group("/products", middlewares.CheckAuth)
	{
		products.GET("/", productcontroller.Index)
		products.GET("/add", productcontroller.Add)
		products.POST("/add", productcontroller.Add)
		products.GET("/detail/:id", productcontroller.Detail)
		products.GET("/edit", productcontroller.Edit)
		products.POST("/edit", productcontroller.Edit)
		products.GET("/delete", productcontroller.Delete)
	}

	log.Println("server running on port 8080")
	r.Run(":8080")
}
