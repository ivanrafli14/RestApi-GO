package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanrafli14/OpenMusicAPI-dicoding/controller/albumcontroller"
	"github.com/ivanrafli14/OpenMusicAPI-dicoding/controller/songcontroller"
	"github.com/ivanrafli14/OpenMusicAPI-dicoding/model"
)

func main() {
	route := gin.Default()
	model.ConnectDatabase()

	route.POST("/albums", albumcontroller.Create)
	route.GET("/albums/:id", albumcontroller.Show)
	route.PUT("/albums/:id", albumcontroller.Update)
	route.DELETE("/albums/:id", albumcontroller.Delete)

	route.POST("/songs", songcontroller.Create)
	route.GET("/songs", songcontroller.Index)
	route.GET("/songs/:id", songcontroller.Show)
	route.PUT("/songs/:id", songcontroller.Update)
	route.DELETE("/songs/:id", songcontroller.Delete)

	route.Run()
}