package main

import (
	bl "buildTool/src/app/base"

	"github.com/gin-gonic/gin"
)

func main() {

	//hl.Connect()
	// web()
	// hl.Build()
	// hl.MyTest()
	// hl.TestPoint()
	bl.TestSlice()

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
