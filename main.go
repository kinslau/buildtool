package main

import (
	m1 "buildTool/src/app/util"
	"io"
	"os"
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
	result, err := m1.ExecCommandResult("rg", "-No", "--column", num, "./圆周率小数点后00000000001到00100000000一共1亿位.txt")
	if err != nil {
		return ""
	}
	return result
}

func web() {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/wx/query", func(c *gin.Context) {
		birthday := c.Query("birthday")

		content := getResult(birthday, "")
		list := strings.Split(content, "\n")
		index := "0"
		if len(list) > 1 {
			index = strings.Split(list[0], ":")[0]
		}

		text := "你的生日" + birthday + "出现在圆周率π的第" + index + "位。(第0位表示在π的前2亿个数字内找不到)"
		c.String(200, text)
	})

	r.Run(":80")

}
