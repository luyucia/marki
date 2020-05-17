package app

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"marki/render"
	"os"
	"path/filepath"
	"strings"

)

//bleve索引的文档：
//http://blevesearch.com/docs/Getting%20Started/

type FileNode struct {
	FileName string     `json:"name"`
	Type     string     `json:"-"`
	Path     string     `json:"-"`
	Url      string     `json:"url"`
	IsDir    bool       `json:"-"`
	Childs   []FileNode `json:"child"`
}

type Document struct {
	Id            string
	Content       string
	ParsedContent string
	RealPath      string
}

func getHashCode(s string) string {

	h := md5.New()
	h.Write([]byte(s))
	md5str := fmt.Sprintf("%x", h.Sum(nil))
	return md5str
}

func handleName(s string) (string, string) {

	ext := filepath.Ext(s)
	name := strings.TrimSuffix(s, ext)

	return name, ext[1:]

}

// 递归生成文件树
func createFileInfoTree(root string) FileNode {

	filelist, err := ioutil.ReadDir(root)
	if err != nil {
		log.Error(err)
	}

	info, err := os.Stat(root)
	if err != nil {
		log.Error(err)
	}
	r := FileNode{}
	r.IsDir = info.IsDir()
	r.FileName = info.Name()
	r.Path = root

	for _, file := range filelist {
		if file.IsDir() == true {
			// 递归遍历
			r.Childs = append(r.Childs, createFileInfoTree(root+"/"+file.Name()))
		} else {
			doc := Document{}
			doc.RealPath = root + "/" + file.Name()

			name, ext := handleName(file.Name())

			_, exist := Config.AllowType[ext]

			if exist {
				leaf := FileNode{}
				leaf.FileName = name
				leaf.Type = ext
				leaf.IsDir = file.IsDir()
				leaf.Path = root + "/" + file.Name()
				leaf.Url = getHashCode(root + "/" + file.Name())
				r.Childs = append(r.Childs, leaf)
				DocumentMap[leaf.Url] = doc
			}

		}
	}

	return r

}

func GenerateMenu() FileNode {

	//遍历目录，获取需要关注的文件 .md的
	docpath := Config.HomeDir
	fileroot := createFileInfoTree(docpath)

	MenuData = fileroot

	return fileroot

}

func GenerateContent(id string) []byte{

	// 先从缓存读取
	rendered,err := Cache.Get(id)
	//命中缓存
	if err==nil{
		log.Info("从缓存读取"+id)
		return rendered
	}
	// 没读到从文件读
	docpath := DocumentMap[id].RealPath

	content, err := ioutil.ReadFile(docpath)
	if err!=nil {
		log.Error(err)
	}
	// 解析markdown
	rendered = render.MarkdownRender(content)
	// 存入缓存
	Cache.Set(id,rendered)


	return rendered

}
