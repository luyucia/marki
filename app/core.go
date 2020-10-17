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
	"time"
)

// bleve索引的文档：
// http://blevesearch.com/docs/Getting%20Started/

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

func getFileNameAndExt(s string) (string, string) {

	ext := filepath.Ext(s)
	name := strings.TrimSuffix(s, ext)

	if ext == "" {
		ext = "."
	}

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

	// 过滤无效目录
	if r.IsDir == true {
		if ignore(r.FileName) {
			return r
		}
	}

	for _, file := range filelist {
		if file.IsDir() == true {
			// 递归遍历
			r.Childs = append(r.Childs, createFileInfoTree(root+"/"+file.Name()))
		} else {
			doc := Document{}
			doc.RealPath = root + "/" + file.Name()

			name, ext := getFileNameAndExt(file.Name())

			_, allow := GConfig.AllowType[ext]

			// 只处理准许的类型
			if allow {
				leaf := FileNode{}
				leaf.FileName = name
				leaf.Type = ext
				leaf.IsDir = file.IsDir()
				leaf.Path = root + "/" + file.Name()
				leaf.Url = getHashCode(root + "/" + file.Name())
				r.Childs = append(r.Childs, leaf)
				GDocumentMap[leaf.Url] = doc
			}

		}
	}

	return r

}

// 递归判断子节点中是否有有效节点
func checkChild(fn FileNode) bool {

	for _, n := range fn.Childs {
		if n.Type != "" {
			return true
		}
		if n.IsDir {
			return checkChild(n)
		}
	}

	return false
}

func cleanFileNode(fn FileNode) FileNode {
	newNode := FileNode{}
	newNode.FileName = fn.FileName
	newNode.IsDir = fn.IsDir

	for _, n := range fn.Childs {
		if n.Type != "" {
			newNode.Childs = append(newNode.Childs, n)
		}
		if n.IsDir {
			if checkChild(n) {
				newNode.Childs = append(newNode.Childs, cleanFileNode(n))
			}

		}

	}

	return newNode

}

// 过滤忽略文件夹
func ignore(dirname string) bool {
	// todo：改成配置化，用hashmap,可让用户指定
	if dirname == ".git" {
		return true
	}

	return false
}

func GenerateMenu() FileNode {
	// 加锁防止 并发调用会有问题
	GMenuGenerateMutex.Lock()
	// 不要频繁生成，让这个函数执行最多每秒1次，获取上次执行时间，如果当前时间-上次时间<1s 则不执行
	lastUpdateTimeDiff := time.Now().Sub(GLastMenuUpdateTime).Seconds()
	if lastUpdateTimeDiff < 1 {
		GMenuGenerateMutex.Unlock()
		return GMenuData
	}
	log.Info("构建菜单")
	docpath := GConfig.HomeDir
	// 遍历目录，获取需要关注的文件 .md的
	// todo:注意此处多次递归 可能影响性能 约2ms后期可以优化
	start_time := time.Now()
	fileroot := createFileInfoTree(docpath)
	fileroot = cleanFileNode(fileroot)

	log.Info("构建目录 耗时:", time.Now().Sub(start_time).Milliseconds(), " ms")

	GMenuData = fileroot
	GLastMenuUpdateTime = time.Now()

	GMenuGenerateMutex.Unlock()
	return fileroot

}

func GenerateContent(id string) ([]byte, string) {

	// 没读到从文件读
	docpath := GDocumentMap[id].RealPath
	// 根据扩展名解析
	_, ext := getFileNameAndExt(GDocumentMap[id].RealPath)

	// 先从缓存读取
	rendered, err := GCache.Get(id)
	// 命中缓存
	if err == nil {
		log.Info("从缓存读取" + id)
		return rendered, ext
	}
	// 没命中缓存
	content, err := ioutil.ReadFile(docpath)
	if err != nil {
		log.Error(err)
	}

	switch ext {
	case "md":
		// 解析markdown
		rendered = render.MarkdownRender(content)
	case "map":
		rendered = render.MindMapRender(content)
	case "http":
		rendered = render.HttpRender(content)
	default:
		rendered = render.MarkdownRender(content)
	}

	// 存入缓存
	GCache.Set(id, rendered)

	return rendered, ext

}
