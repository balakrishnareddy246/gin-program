package main

import (
	"math/rand"
	"time"

	""
)

func main() {

	rand.Seed(time.Now().UnixNano())
	router := gin.Default()

	router.Static("/img", "./img")
	router.GET("/", mainHandler)
	router.GET("/Kitty", KittyHandler)
	router.GET("/pup", pupHandler)
	router.Run(":8081")
}
