package main

import (
	m1 "buildTool/src/app/util"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var logger = m1.GetLogger()

func main() {

	// hl.Connect()
	// hl.Build()
	// MyTest()
	// hl.TestPoint()
	// bl.TestSlice()

	web()

}

func testLog() {

	for {
		time.Sleep(30 * time.Millisecond)
		logger.Info("----lucky number is ", rand.Intn(10000000))

	}

}

func web() {

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		go testLog()
		c.JSON(200, gin.H{
			"message": "pong",
			"item":    "",
		})
	})
	r.Run()

}
