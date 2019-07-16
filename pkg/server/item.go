package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ledger/pkg/database"
	"ledger/pkg/models"
)

func parseUint(str string) uint32 {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("err=%s\n", err)
	}
	return uint32(val)
}

func parseFloat(str string) float32 {
	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		fmt.Printf("err=%s\n", err)
	}
	return float32(val)
}

func GetItems(c *gin.Context) (items []models.Item) {
	db := database.GetDB()
	if err := db.Preload("Category").Find(&items).Error; err != nil {
		fmt.Printf("err=%s\n", err)
	}
	return
}

func ItemIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "item_index.tmpl", gin.H{
		"items": GetItems(c),
	})
}

func ItemNew(c *gin.Context) {
	c.HTML(http.StatusOK, "item_new.tmpl", gin.H{
		"item": models.Item{},
		"categories": GetCategories(c),
	})
}

func ItemUpdate(c *gin.Context) {
	c.HTML(http.StatusOK, "item_new.tmpl", gin.H{
		"item": GetApiItem(c),
		"categories": GetCategories(c),
	})
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

func GetApiItem(c *gin.Context) (item models.Item) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		fmt.Printf("err=%s\n", err)
		return
	}
	db := database.GetDB()
	if err := db.Preload("Category").
		Where("items.id = ?", id).
		Find(&item).Error; err != nil {
		fmt.Printf("err=%s\n", err)
	}
	return
}

func PostApiItem(c *gin.Context) {
	var item models.Item
	if err := c.Request.ParseForm(); err != nil {
		fmt.Printf("err=%s\n", err)
		return
	}
	data := c.Request.PostForm

	item.Name = data["name"][0]
	item.Amount = parseFloat(data["amount"][0])
	item.CategoryId = parseUint(data["category_id"][0])
	db := database.GetDB()
	if err := db.Create(&item).Error; err != nil {
		fmt.Printf("err=%s\n", err)
	}
	c.Redirect(http.StatusMovedPermanently, "/item")
}

func PutApiItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		fmt.Printf("err=%s\n", err)
	}
	db := database.GetDB()
	if err := db.Update(&item).Error; err != nil {
		fmt.Printf("err=%s\n", err)
	}
	c.JSON(http.StatusOK, item)
}

func DeleteApiItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		fmt.Printf("err=%s\n", err)
		return
	}
	db := database.GetDB()
	if err := db.Where("id = ?", id).
		Delete(&models.Item{}).Error; err != nil {
		fmt.Printf("err=%s\n", err)
	}
	c.Redirect(http.StatusMovedPermanently, "/item")
}
