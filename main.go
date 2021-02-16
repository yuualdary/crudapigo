package main

import (
	"Assigment2/config"
	"Assigment2/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/gorilla/mux"
)

func main() {

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	var router = gin.Default()

	//route

	router.PUT("/update/:id", inDB.UpdateOrder)
	router.POST("/create", inDB.CreateOrder)

	router.Run(":3000")

}
