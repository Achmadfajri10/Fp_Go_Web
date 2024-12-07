package main

import (
	"Fp_Go_Web/config"
	homecontroller "Fp_Go_Web/controllers"
	"Fp_Go_Web/controllers/categorycontroller"
	"Fp_Go_Web/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//1.Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//2.Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	//3.Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
