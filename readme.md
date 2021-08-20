
### 打包构建

#### Windows下编译Mac, Linux平台的64位可执行程序：
```sh

$ SET GOOS=linux
$ SET GOOS=windows
$ SET GOARCH=amd64
$ go build main.go

```