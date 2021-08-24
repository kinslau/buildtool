package main

import (
	m1 "buildTool/src/app/util"
	"encoding/xml"
	"fmt"
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

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.POST("/wx/token", func(c *gin.Context) {

		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}

		xmlStr := string(data)

		result := WXRequest{}
		xml.Unmarshal([]byte(xmlStr), &result)

		logger.Info("ToUserName:", result.ToUserName.Text)
		logger.Info("FromUserName:", result.FromUserName.Text)
		logger.Info("content:", result.Content.Text)

		content := getResult(result.Content.Text, "")
		list := strings.Split(content, "\n")
		index := "0"
		if len(list) > 1 {
			index = strings.Split(list[0], ":")[0]
		}

		text := "你的生日" + result.Content.Text + "出现在圆周率π的第" + index + "位。(第0位表示在π的前2亿个数字内找不到)"

		msg := WX{
			ToUserName:   CDATA{result.FromUserName.Text},
			FromUserName: CDATA{result.ToUserName.Text},
			CreateTime:   int(time.Now().Unix()),
			MsgType:      CDATA{"text"},
			Content:      CDATA{text},
		}
		c.XML(200, msg)

	})

	r.GET("/wx/token", func(c *gin.Context) {
		echostr := c.Query("echostr")
		c.String(200, echostr)
	})

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

type WX struct {
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int
	MsgType      CDATA
	Content      CDATA
	XMLName      xml.Name `xml:"xml"`
}
type CDATA struct {
	Text string `xml:",cdata"`
}

type WXRequest struct {
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int
	MsgType      CDATA
	Content      CDATA
	MsgId        int64
	XMLName      xml.Name `xml:"xml"`
}
