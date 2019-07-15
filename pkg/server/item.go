package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"ledger/pkg/database"
	"ledger/pkg/models"
)

func GetItems(c *gin.Context) (items []models.Item) {
	db := database.GetDB()
	if err := db.Preload("Category").Find(&items).Error; err != nil {
		fmt.Printf("err=%s\n", err)
	}
	return
}

func GetApiItems(c *gin.Context) {
	var items []models.Item
	db := database.GetDB()
	if err := db.Preload("Category").Find(&items).Error; err != nil {
		fmt.Printf("err=%s\n", err)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

func PostApiItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		fmt.Printf("err=%s\n", err)
	}
	db := database.GetDB()
	if err := db.Create(&item).Error; err != nil {
		fmt.Printf("err=%s\n", err)
	}
	c.JSON(http.StatusOK, item)
}
