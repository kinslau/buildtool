package main

import (
	hl "buildTool/src/app/util"

	"github.com/gin-gonic/gin"
)

func main() {
	web()
	hl.Build()
}

func web() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
			"item":    "",
		})
	})
	r.Run()
}
