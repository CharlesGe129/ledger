package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ledger/pkg/database"
	"ledger/pkg/models"
	"net/http"
)

func GetItems(c *gin.Context) (items []models.Item) {
	db := database.GetDB()
	if err := db.Table("items").Joins("JOIN categories ON items.category_id=categories.id").Find(&items).Error; err != nil {
		fmt.Printf("err=%s", err)
	}
	return
}

func GetApiItems(c *gin.Context) {
	var items []models.Item
	db := database.GetDB()
	if err := db.Table("items").Joins("JOIN categories ON items.category_id=categories.id").Find(&items).Error; err != nil {
		fmt.Printf("err=%s", err)
	} else {
		c.JSON(http.StatusOK, items)
	}
}
