package render

// render 层作用是，渲染原始数据 产生目标数据，
// 目标数据为html格式用来前端展示

// 目前支持markdown
// todo：支持api
// todo：支持脚本语言运行

import (
	//"github.com/russross/blackfriday/v2"
	"github.com/russross/blackfriday"
)

func MarkdownRender(s []byte) []byte {

	// 参考 https://github.com/russross/blackfriday-tool/blob/master/main.go

	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	var renderer blackfriday.Renderer
	renderer = blackfriday.HtmlRenderer(0, "", "")

	html := blackfriday.Markdown(s, renderer, extensions)

	// v2.0 不好使 markdown语法解析为html
	//unsafe := blackfriday.Run(s, blackfriday.WithExtensions(blackfriday.Tables),blackfriday.WithExtensions(blackfriday.FencedCode))
	//html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	return html

}
