package controllers

import (
	"Assigment2/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//to find product
// func (idb *InDB) GetOrder(c *gin.Context) {

// 	var (
// 		orders structs.Orders
// 		items  structs.Items

// 		result gin.H
// 	)

// 	id := c.Param("id") //digunakan untuk get
// 	var x = db.Model(&items).Select("*").Joins("left join orders on orders.order_id = items.order_id").Scan(&result{})

// 	var err = idb.DB.Where("Order_id = ?", id).First(&orders).Error

// 	if err != nil {
// 		result = gin.H{
// 			"result": err.Error(),
// 			"count":  0,
// 		}
// 	} else {
// 		result = gin.H{
// 			"result": orders,
// 			"count":  1,
// 		}
// 	}
// 	c.JSON(http.StatusOK, result)

// }

//createproduct
func (idb *InDB) CreateOrder(c *gin.Context) {

	var (
		order  structs.Orders
		item   structs.Items
		result gin.H
	)

	// id := c.PostForm("Id")
	customername := c.PostForm("costumer_name")
	description := c.PostForm("description")
	item_code := c.PostForm("item_code")
	quantity := c.PostForm("quantity")

	order.Customer_name = customername
	item.Description = description
	item.Item_code, _ = strconv.Atoi(item_code)
	item.Quantity, _ = strconv.Atoi(quantity)

	idb.DB.Create(&order)
	idb.DB.Create(&item)

	result = gin.H{
		"result": order,
		"item":   item,
	}
	c.JSON(http.StatusOK, result)
}

//update
func (idb *InDB) UpdateOrder(c *gin.Context) {

	id := c.Params.ByName("id")
	customername := c.PostForm("costumer_name")
	description := c.PostForm("description")
	item_code := c.PostForm("item_code")
	quantity := c.PostForm("quantity")

	var (
		order    structs.Orders
		newOrder structs.Orders
		item     structs.Items
		newItem  structs.Items

		result gin.H
	)

	err := idb.DB.First(&item, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newOrder.Customer_name = customername
	newItem.Description = description
	newItem.Item_code, _ = strconv.Atoi(item_code)
	newItem.Quantity, _ = strconv.Atoi(quantity)

	err = idb.DB.Model(&order).Updates(newOrder).Error

	if err != nil {
		result = gin.H{
			"result": "Data Not Found",
		}
	}
	c.JSON(http.StatusOK, result)

}

func (idb *InDB) DeleteOrder(c *gin.Context) {

	var (
		// order  structs.Orders
		item   structs.Items
		result gin.H
	)

	id := c.Param("id")
	it := idb.DB.Where("item_id = ?", id).First(&item).Error

	if it != nil {
		result = gin.H{
			"Result": "Data Not Found",
		}
	}
	it = idb.DB.Delete(&item).Error
	if it != nil {
		result = gin.H{
			"result": "Delete Failed",
		}
	} else {
		result = gin.H{
			"result": "Success Delete",
		}
	}
	c.JSON(http.StatusOK, result)
}
