package productcontroller

import (
	"Fp_Go_Web/entities"
	"Fp_Go_Web/models/categorymodel"
	"Fp_Go_Web/models/productmodel"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	products := productmodel.GetAll()
	c.HTML(http.StatusOK, "productindex.html", gin.H{
		"products": products,
	})
}

func Detail(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid product ID")
		return
	}

	product := productmodel.Detail(id)
	if product.Id == 0 {
		c.String(http.StatusNotFound, "Product not found")
		return
	}

	c.HTML(http.StatusOK, "productdetail.html", gin.H{
		"product": product,
	})
}

func Add(c *gin.Context) {
	if c.Request.Method == "GET" {
		categories := categorymodel.GetAll()
		c.HTML(http.StatusOK, "productcreate.html", gin.H{
			"categories":   categories,
			"noCategories": len(categories) == 0,
		})
		return
	}

	if c.Request.Method == "POST" {
		categories := categorymodel.GetAll()
		if len(categories) == 0 {
			c.Redirect(http.StatusSeeOther, "/products/add?error=no_categories")
			return
		}

		categoryId, err := strconv.Atoi(c.PostForm("category_id"))
		if err != nil || categoryId == 0 {
			c.Redirect(http.StatusSeeOther, "/products/add?error=invalid_category")
			return
		}

		stock, err := strconv.Atoi(c.PostForm("stock"))
		if err != nil || stock < 0 {
			c.Redirect(http.StatusSeeOther, "/products/add?error=invalid_stock")
			return
		}

		product := entities.Product{
			Name:        c.PostForm("name"),
			Category:    entities.Category{Id: uint(categoryId)},
			Stock:       stock,
			Description: c.PostForm("description"),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if ok := productmodel.Create(product); !ok {
			c.Redirect(http.StatusSeeOther, "/products/add?error=create_failed")
			return
		}

		c.Redirect(http.StatusSeeOther, "/products")
	}

}

func Edit(c *gin.Context) {
	if c.Request.Method == "GET" {
		idString := c.Query("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid product ID")
			return
		}

		product := productmodel.Detail(id)
		categories := categorymodel.GetAll()
		c.HTML(http.StatusOK, "productedit.html", gin.H{
			"categories": categories,
			"product":    product,
		})
		return
	}

	if c.Request.Method == "POST" {
		idString := c.PostForm("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid product ID")
			return
		}

		categoryId, err := strconv.Atoi(c.PostForm("category_id"))
		if err != nil || categoryId == 0 {
			c.Redirect(http.StatusSeeOther, "/products/edit?id="+idString+"&error=invalid_category")
			return
		}

		stock, err := strconv.Atoi(c.PostForm("stock"))
		if err != nil || stock < 0 {
			c.Redirect(http.StatusSeeOther, "/products/edit?id="+idString+"&error=invalid_stock")
			return
		}

		product := entities.Product{
			Id: 		uint(id),
			Name:        c.PostForm("name"),
			Category:    entities.Category{Id: uint(categoryId)},
			Stock:       stock,
			Description: c.PostForm("description"),
			UpdatedAt:   time.Now(),
		}

		if ok := productmodel.Update(id, product); !ok {
			fmt.Printf("Failed to update product with ID %s\n", ok)
			c.Redirect(http.StatusSeeOther, "/products/edit?id="+idString+"&error=update_failed")
			return
		}

		c.Redirect(http.StatusSeeOther, "/products")
	}
}

func Delete(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := productmodel.Delete(id); err != nil {
		c.String(http.StatusInternalServerError, "Error deleting product")
		return
	}

	c.Redirect(http.StatusSeeOther, "/products")
}
