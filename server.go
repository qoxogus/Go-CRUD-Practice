package main

import (
	"goCRUD/config"
	"goCRUD/controller"
	"goCRUD/database"

	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	config.InitConfig()
	database.Connect()

	r := gin.Default()

	r.GET("/", controller.Home)
	r.GET("/test", controller.Test)

	r.POST("/create", controller.Create) //insert 로 하겠음
	r.POST("/read", controller.Read)     // select
	r.POST("/update", controller.Update) // update
	r.POST("/delete", controller.Delete) // delete

	r.Run(":0809")
}
