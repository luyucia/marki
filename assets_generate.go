//+build ignore

package main

import (
	"fmt"
	"github.com/shurcooL/vfsgen"
	"net/http"
)

//go:generate go run assets_generate.go
// 生成静态文件，在执行go generate命令时执行此函数
func main() {

	var fs http.FileSystem = http.Dir("H:/marki/web/dist")
	err := vfsgen.Generate(fs, vfsgen.Options{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("生成静态文件")
}
