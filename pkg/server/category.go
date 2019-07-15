package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ledger/pkg/database"
	"ledger/pkg/models"
)

func GetCategories(c *gin.Context) (cats []models.Category) {
	db := database.GetDB()
	if err := db.Find(&cats).Error; err != nil {
		fmt.Printf("err=%s", err)
	}
	return
}
