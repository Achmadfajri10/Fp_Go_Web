package productcontroller

import (
	"Fp_Go_Web/entities"
	"Fp_Go_Web/models/categorymodel"
	"Fp_Go_Web/models/productmodel"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	products, err := productmodel.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve products")
		return
	}
	c.HTML(http.StatusOK, "productindex.html", gin.H{"products": products})
}

func Detail(c *gin.Context) {
	idString := c.Query("id")
	if idString == "" {
		c.String(http.StatusBadRequest, "Missing product ID")
		return
	}

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := productmodel.Detail(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.String(http.StatusNotFound, "Product not found")
		} else {
			c.String(http.StatusInternalServerError, "Failed to retrieve product")
		}
		return
	}
	c.HTML(http.StatusOK, "productdetail.html", gin.H{"product": product})
}

func Add(c *gin.Context) {
	if c.Request.Method == "GET" {
		categories, err := categorymodel.GetAll()
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to retrieve categories")
			return
		}
		c.HTML(http.StatusOK, "productcreate.html", gin.H{
			"categories":   categories,
			"noCategories": len(categories) == 0,
		})
		return
	}

	if c.Request.Method == "POST" {
		categories, err := categorymodel.GetAll()
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to retrieve categories")
			return
		}
		if len(categories) == 0 {
			c.Redirect(http.StatusSeeOther, "/products/add?error=no_categories")
			return
		}

		categoryId, err := strconv.ParseUint(c.PostForm("category_id"), 10, 64)
		if err != nil || categoryId == 0 {
			c.Redirect(http.StatusSeeOther, "/products/add?error=invalid_category")
			return
		}

		stock, err := strconv.ParseInt(c.PostForm("stock"), 10, 32)
		if err != nil || stock < 0 {
			c.Redirect(http.StatusSeeOther, "/products/add?error=invalid_stock")
			return
		}

		product := entities.Product{
			Name:        c.PostForm("name"),
			CategoryID:  uint(categoryId),
			Stock:       int(stock),
			Description: c.PostForm("description"),
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
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid product ID")
			return
		}

		product, err := productmodel.Detail(uint(id))
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to retrieve product")
			return
		}
		categories, err := categorymodel.GetAll()
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to retrieve categories")
			return
		}
		c.HTML(http.StatusOK, "productedit.html", gin.H{
			"categories": categories,
			"product":    product,
		})
		return
	}

	if c.Request.Method == "POST" {
		idString := c.PostForm("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid product ID")
			return
		}

		categoryId, err := strconv.ParseUint(c.PostForm("category_id"), 10, 64)
		if err != nil || categoryId == 0 {
			c.Redirect(http.StatusSeeOther, "/products/edit?id="+idString+"&error=invalid_category")
			return
		}

		stock, err := strconv.ParseInt(c.PostForm("stock"), 10, 32)
		if err != nil || stock < 0 {
			c.Redirect(http.StatusSeeOther, "/products/edit?id="+idString+"&error=invalid_stock")
			return
		}

		product := entities.Product{
			Name:        c.PostForm("name"),
			CategoryID:  uint(categoryId),
			Stock:       int(stock),
			Description: c.PostForm("description"),
		}

		err = productmodel.Update(uint(id), product)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/products/edit?id="+idString+"&error=update_failed")
			return
		}

		c.Redirect(http.StatusSeeOther, "/products")
	}
}

func Delete(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = productmodel.Delete(uint(id))
	if err != nil {
		c.String(http.StatusInternalServerError, "Error deleting product")
		return
	}

	c.Redirect(http.StatusSeeOther, "/products")
}
