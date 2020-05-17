package render
// render 层作用是，渲染原始数据 产生目标数据，
// 目标数据为html格式用来前端展示

// 目前支持markdown
// todo：支持api
// todo：支持脚本语言运行

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)


func MarkdownRender(s []byte) []byte {

	// markdown语法解析为html
	unsafe := blackfriday.Run(s)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	return html

}
