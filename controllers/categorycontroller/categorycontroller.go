package categorycontroller

import (
	"Fp_Go_Web/entities"
	"Fp_Go_Web/models/categorymodel"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	categories, err := categorymodel.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve categories")
		return
	}
	c.HTML(http.StatusOK, "categoryindex.html", gin.H{
		"categories": categories,
	})
}

func Add(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "categorycreate.html", nil)
		return
	}

	if c.Request.Method == "POST" {
		var category entities.Category
		category.Name = c.PostForm("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Create(category); !ok {
			c.HTML(http.StatusInternalServerError, "categorycreate.html", gin.H{"error": "Failed to create category"})
			return
		}

		c.Redirect(http.StatusSeeOther, "/categories/")
	}

}

func Edit(c *gin.Context) {

	if c.Request.Method == "GET" {
		idString := c.Query("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}

		category, err := categorymodel.Detail(uint(id))
		if err == nil {
			c.HTML(http.StatusOK, "categoryedit.html", gin.H{
				"category": category,
			})
		}
		return
	}

	if c.Request.Method == "POST" {
		var category entities.Category

		idString := c.PostForm("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}

		category.Name = c.PostForm("name")
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Update(uint(id), category); !ok {
			c.Redirect(http.StatusSeeOther, "/categories/edit?id="+idString)
			return
		}

		c.Redirect(http.StatusSeeOther, "/categories/")
	}
}

func Delete(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := categorymodel.Delete(uint(id)); err != nil {
		c.Redirect(http.StatusSeeOther, "/categories")
		return
	}

	c.Redirect(http.StatusSeeOther, "/categories")
}
