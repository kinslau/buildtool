package main

import (
	m1 "buildTool/src/app/util"
	"strings"

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

func getResult(num string, path string) string {
	result, err := m1.ExecCommandResult("rg", "-No", "--column", num, "./圆周率小数点后24900000001到25000000000一共1亿位.txt")
	if err != nil {
		return ""
	}
	return result
}

func web() {

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		args := c.Query("args")
		result := getResult(args, "")
		logger.Info(result)
		var list []string
		list = strings.Split(result, "\n")

		c.JSON(200, gin.H{
			"message": list,
		})
	})
	r.Run(":80")

}
