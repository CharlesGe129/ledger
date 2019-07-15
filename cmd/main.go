package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"ledger/pkg/database"
	"ledger/pkg/server"
	"net/http"
)

var db *gorm.DB

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello World",
	})
}

func itemIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "item_index.tmpl", gin.H{
		"title": "Expenses",
		"items": server.GetItems(c),
	})
}

func catIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "cat_index.tmpl", gin.H{
		"title": "Categories",
		"cats": server.GetCategories(c),
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	database.GetDB()

	r.GET("/", index)
	r.GET("/item", itemIndex)

	r.GET("/api/item", server.GetApiItems)
	r.POST("/api/item", server.PostApiItem)
	r.PUT("/api/item", server.PutApiItem)

	//r.GET("/api/category", server.GetCategories)
	//r.POST("/api/category", server.CreateCategories)

	r.Run(":9000")
}
