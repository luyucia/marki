package app

import (
	"github.com/blevesearch/bleve"
	"github.com/labstack/gommon/log"
)

// 执行索引操作
func indexing()  {
	log.Info("indexing...")
	//根据目录列表，索引每个文件的内容
	//MenuData.Childs
	//for k,v := range DocumentMap{
	//	DocIndex.Index(k,GenerateContent(k))
	//	log.Info("index the doc :"+v.RealPath)
	//}
	rs := query("doraemon")
	log.Info(rs)

}

func query(keyword string) *bleve.SearchResult {
	query := bleve.NewMatchQuery(keyword)
	search := bleve.NewSearchRequest(query)

	searchResult,err := DocIndex.Search(search)

	if err!= nil {
		log.Error(err)
	}
	return searchResult

}
