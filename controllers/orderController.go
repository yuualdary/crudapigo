package controllers

import (
	"Assigment2/config"
	"Assigment2/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCurrentOrderById(c *gin.Context) {
	var (
		id    = c.Params.ByName("id")
		order []structs.Order
	)

	if err := config.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, order)

}

func CreateOrder(c *gin.Context) {

	var (
		order structs.Order
	)

	c.ShouldBindJSON(&order)

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error Item": "Record can't create!",
		})
	}

	c.JSON(http.StatusOK, order)
}

//update

func UpdateOrder(c *gin.Context) {
	id := c.Params.ByName("id")

	var (
		order structs.Order
		// status structs.Status
	)

	// if err := config.DB.Select("status_id").Where("statusname = ?", "Updated").First(&status).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }
	// fmt.Println(status)

	if err := config.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&order)

	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't update!"})
		return
	}

	c.JSON(http.StatusOK, order)

}

func DeleteStudentById(c *gin.Context) {
	id := c.Params.ByName("id")

	var (
		item  structs.Item
		order structs.Order
	)

	if err := config.DB.Where("Order_id = ?", id).Delete(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	if err := config.DB.Where("Order_id = ?", id).Delete(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})

	}

	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}
