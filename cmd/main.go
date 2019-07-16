package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"ledger/pkg/database"
	"ledger/pkg/models"
	"ledger/pkg/server"
	"html/template"
	"net/http"
	"time"
)

var db *gorm.DB

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello World",
	})
}

func itemIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "item_index.tmpl", gin.H{
		"items": server.GetItems(c),
	})
}

func itemNew(c *gin.Context) {
	c.HTML(http.StatusOK, "item_new.tmpl", gin.H{
		"item": models.Item{},
		"categories": server.GetCategories(c),
	})
}

func itemUpdate(c *gin.Context) {
	c.HTML(http.StatusOK, "item_new.tmpl", gin.H{
		"item": server.GetApiItem(c),
		"categories": server.GetCategories(c),
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
	r.GET("/item", itemIndex)
	r.GET("/item/new", itemNew)
	r.PUT("/item/update", itemUpdate)

	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/item", server.GetApiItems)
		apiRoutes.POST("/item", server.PostApiItem)
		apiRoutes.PUT("/item", server.PutApiItem)

		apiRoutes.GET("/cats", server.GetCats)
	}
	//r.POST("/api/category", server.CreateCategories)

	r.Run(":9099")
}
