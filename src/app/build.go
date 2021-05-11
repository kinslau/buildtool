package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func main() {

	config := InitConfig("config.properties")
	sourceCommit := config["sourceCommit"]
	targetCommit := config["targetCommit"]
	baseDir := config["baseDir"]
	targetDir := config["targetDir"]
	appName := config["appName"]

	if sourceCommit == "" || targetCommit == "" || baseDir == "" {
		fmt.Println("读取配置文件失败")
		panic("读取配置文件失败")
	}

	deleteFiles(appName)
	syncGit(baseDir)
	mvnPackage(baseDir)

	var diffFiles = getSourceDiffFiles(appName, sourceCommit, targetCommit, baseDir)
	var targetFiles = getTargetFiles(diffFiles, targetDir)
	createFile(appName+"_target.txt", targetFiles)
	createTargetFiles(appName, targetFiles, targetDir)

}

func syncGit(baseDir string) {
	{
		log.Println("------GIT仓库检查------")
		cmd := exec.Command("git", "status")
		cmd.Dir = baseDir
		result, err := cmd.Output()
		if err != nil {
			text := err.Error()
			fmt.Println(err)
			panic("GIT仓库检查失败" + text)
		}
		fmt.Println(string(result))
	}

	{
		log.Println("------切换分支------")
		cmd := exec.Command("git", "checkout", "lb-pkg")
		cmd.Dir = baseDir
		result, err := cmd.Output()
		if err != nil {
			text := err.Error()
			fmt.Println(err)
			panic("切换分支失败" + text)
		}
		fmt.Println(string(result))
	}

	{
		log.Println("------拉取最新分支------")
		cmd := exec.Command("git", "fetch", "--all")
		cmd.Dir = baseDir
		result, err := cmd.Output()
		if err != nil {
			text := err.Error()
			fmt.Println(err)
			panic("拉取最新分支失败" + text)
		}
		fmt.Println(string(result))
	}

	{
		log.Println("------同步代码------")
		cmd := exec.Command("git", "pull", "origin", "lb-pkg")
		cmd.Dir = baseDir
		result, err := cmd.Output()
		if err != nil {
			text := err.Error()
			fmt.Println(err)
			panic("同步代码失败" + text)
		}

		fmt.Println(string(result))
	}

}

func mvnPackage(baseDir string) {
	log.Println("------MAVEN开始构建------")
	cmd := exec.Command("mvn", "clean", "install", "-Psit1")
	cmd.Dir = baseDir
	result, err := cmd.Output()
	if err != nil {
		text := err.Error()
		fmt.Println(err)
		panic("MAVEN构建失败" + text)
	}

	garbledStr := ConvertByte2String(result, GB18030)
	fmt.Println(garbledStr)

	log.Println("------MAVEN构建完成------")
}

func getSourceDiffFiles(appName string, sourceCommit string, targetCommit string, baseDir string) []string {
	var diffFiles = make([]string, 0)
	cmd := exec.Command("git", "diff", sourceCommit, targetCommit, "--name-only")
	cmd.Dir = baseDir
	result, err := cmd.Output()
	if err != nil {
		text := err.Error()
		fmt.Println(err)

		panic("查找GIT差异文件失败" + text)
	}

	var metrics bytes.Buffer
	metrics.Write(result)
	scanner := bufio.NewScanner(&metrics)

	for scanner.Scan() {
		text := scanner.Text()
		diffFiles = append(diffFiles, text)
	}

	_ = ioutil.WriteFile("_source.txt", result, 0777)
	fmt.Println("共查找到个差异文件写入文件source.txt", len(diffFiles))
	return diffFiles
}

func getTargetFiles(diffFiles []string, targetDir string) []string {
	var targetFiles = make([]string, 0)
	for _, f := range diffFiles {

		var text string
		var classes_flag = strings.HasPrefix(f, "src/main")
		var resource_flag_webapp = strings.HasPrefix(f, "src/main/resources")
		var resource_flag_webcontent = strings.HasPrefix(f, "WebContent")
		var web_flag = strings.HasPrefix(f, "src/main/webapp")
		var jar_flg = strings.HasSuffix(f, ".jar")

		if classes_flag && !resource_flag_webapp && !resource_flag_webcontent {

			strs := strings.Split(f, "/")
			strs = append(strs[:0], strs[3:]...)

			text = strings.Join(strs, "/")
			text = "WEB-INF/classes/" + text
			text = strings.Replace(text, ".java", ".class", 1)

		}

		if resource_flag_webapp && !jar_flg {
			text = strings.Replace(f, "src/main/resources", "WEB-INF/classes", 1)
		}

		if web_flag && !jar_flg {
			text = strings.Replace(f, "src/main/webapp", "", 1)
		}
		if resource_flag_webcontent && !jar_flg {
			text = strings.Replace(f, "WebContent/", "", 1)
		}

		if jar_flg {
			text = strings.Replace(f, "src/main/webapp", "", 1)
		}

		targetFiles = append(targetFiles, text)

	}

	return targetFiles
}

func createTargetFiles(appName string, targetFiles []string, targetDir string) {
	for _, text := range targetFiles {
		pathText := text
		pathText = strings.Replace(pathText, "/", `\`, -1)
		srcFile := targetDir + `\` + pathText

		targetFile := GetCurrentPath() + `\target\` + appName + `\` + pathText
		CopyFile(srcFile, targetFile)
	}

}

func deleteFiles(appName string) {
	fmt.Println("-----清除文件------")
	{
		targetPath := GetCurrentPath() + `\target\`
		err := os.RemoveAll(targetPath)
		if err != nil {
			panic("删除失败")
		}
	}

}

/**
创建文件
*/
func createFile(filename string, contents []string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	for _, v := range contents {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

/**
初始化配置
*/
func InitConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}

//判断文件或目录是否存在
func GetFileInfo(src string) os.FileInfo {
	if fileInfo, e := os.Stat(src); e != nil {
		if os.IsNotExist(e) {
			return nil
		}
		return nil
	} else {
		return fileInfo
	}
}

//拷贝文件
func CopyFile(src, dst string) bool {
	if len(src) == 0 || len(dst) == 0 {
		return false
	}
	srcFile, e := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if e != nil {
		println("copyfile", e)
		return false
	}
	defer srcFile.Close()

	dst = strings.Replace(dst, "\\", "/", -1)
	dstPathArr := strings.Split(dst, "/")
	dstPathArr = dstPathArr[0 : len(dstPathArr)-1]
	dstPath := strings.Join(dstPathArr, "/")

	dstFileInfo := GetFileInfo(dstPath)
	if dstFileInfo == nil {
		if e := os.MkdirAll(dstPath, os.ModePerm); e != nil {
			println("copyfile", e)
			return false
		}
	}
	//这里要把O_TRUNC 加上，否则会出现新旧文件内容出现重叠现象
	dstFile, e := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if e != nil {
		println("copyfile", e)
		return false
	}
	defer dstFile.Close()
	//fileInfo, e := srcFile.Stat()
	//fileInfo.Size() > 1024
	//byteBuffer := make([]byte, 10)
	if _, e := io.Copy(dstFile, srcFile); e != nil {
		println("copyfile", e)
		return false
	} else {
		return true
	}

}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
