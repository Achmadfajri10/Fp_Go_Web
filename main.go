package main

import (
	"Fp_Go_Web/config"
	homecontroller "Fp_Go_Web/controllers"
	"Fp_Go_Web/controllers/categorycontroller"
	"Fp_Go_Web/controllers/productcontroller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB() 
	r := gin.Default() 

	r.LoadHTMLGlob("views/**/*.html")

	//1.Homepage
	r.GET("/", homecontroller.Welcome)

	//2.Categories
	categories := r.Group("/categories")
	{
		categories.GET("/", categorycontroller.Index)
		categories.GET("/add", categorycontroller.Add)
		categories.POST("/add", categorycontroller.Add)
		categories.GET("/edit", categorycontroller.Edit)
		categories.POST("/edit", categorycontroller.Edit)
		categories.GET("/delete", categorycontroller.Delete)
	}

	//3.Products
	products := r.Group("/products")
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
