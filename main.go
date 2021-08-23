package main

import (
	m1 "buildTool/src/app/util"
	"encoding/json"
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
		list := strings.Split(result, "\n")

		c.JSON(200, gin.H{
			"message": list,
		})
	})

	r.POST("/wx/token", func(c *gin.Context) {

		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}

		xmlStr := string(data)
		logger.Info("请求body内容为:", xmlStr)

		result := WXRequest{}
		xml.Unmarshal([]byte(xmlStr), &result)

		num := result.Content.Text
		content := getResult(num, "")

		msg := WX{
			ToUserName:   CDATA{result.FromUserName.Text},
			FromUserName: CDATA{result.ToUserName.Text},
			CreateTime:   int(time.Now().Unix()),
			MsgType:      CDATA{"text"},
			Content:      CDATA{content},
		}

		c.XML(200, msg)

	})

	r.GET("/wx/token", func(c *gin.Context) {
		echostr := c.Query("echostr")
		c.String(200, echostr)
	})

	r.Run(":80")

}

func getAccess_token() string {

	appid := "wx6a83652b8ce66170"
	APPSECRET := "e4a6561ab0910b4e2c0759e329e244d7"
	result := m1.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + APPSECRET)

	// an arbitrary json string
	var jsonMap map[string]interface{}

	json.Unmarshal([]byte(result), &jsonMap)

	logger.Info(jsonMap)
	// prints: map[foo:map[baz:[1 2 3]]]

	return result
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
