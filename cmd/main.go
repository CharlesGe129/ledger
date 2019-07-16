package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"ledger/pkg/database"
	"ledger/pkg/server"
)

var db *gorm.DB

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello World",
	})
}

func catIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "cat_index.tmpl", gin.H{
		"title": "Categories",
		"cats": server.GetCategories(c),
	})
}

func formatAsDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLGlob("templates/*")

	database.GetDB()

	r.GET("/", index)
	r.GET("/item", server.ItemIndex)
	r.GET("/item/new", server.ItemNew)
	r.GET("/item/edit/:id", server.ItemUpdate)

	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/item", server.GetApiItems)
		apiRoutes.POST("/item", server.PostApiItem)
		apiRoutes.PUT("/item", server.PutApiItem)
		apiRoutes.GET("/item/delete/:id", server.DeleteApiItem)

		apiRoutes.GET("/cats", server.GetCats)
	}
	//r.POST("/api/category", server.CreateCategories)

	r.Run(":9099")
}
