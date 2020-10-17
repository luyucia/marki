package render

import (
	"io/ioutil"
	"log"
	"testing"
)

type MindNode struct {
	Content string
	Child   []MindNode
}

func TestParse(t *testing.T) {
	data, err := ioutil.ReadFile("mindmap_test_data.txt")
	if err != nil {
		t.Error(err)
		return
	}

	datas := string(data)

	for _, c := range datas {
		log.Print(c)
		log.Print(string(c))

		// 	读取字符到节点名
		//  读取到换行，创建节点，清空节点名，记录父节点指针
		//  读取到tab 继续直到没有tab为止，创建子节点，读取字符串到节点名，读取到换行，赋值节点名，插入父节点的child中，从新开始
	}

	s := MindMapRender(data)

	t.Log(string(s))
}
