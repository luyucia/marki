package app

import (
	"github.com/allegro/bigcache"
	"github.com/blevesearch/bleve"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/gommon/log"
	"time"
)

// 应用配置
type AppConfig struct {
	Port      int
	Host      string
	HomeDir   string
	CacheDir  string
	IndexDir  string
	AllowType map[string]bool
}

// 全局配置
var Config = AppConfig{}
var MenuData FileNode
var DocumentMap map[string]Document
var Cache *bigcache.BigCache
var watcher *fsnotify.Watcher
var DocIndex bleve.Index

//初始化
func init() {
	// 初始化默认参数
	Config.Host = "0.0.0.0"
	Config.Port = 80
	Config.HomeDir = "h:\\mydocs"
	Config.CacheDir = "./cache"
	Config.IndexDir = "./index.bleve"
	Config.AllowType = map[string]bool{}
	Config.AllowType["md"] = true
	Config.AllowType["http"] = true

	// 初始化目录结构
	DocumentMap = map[string]Document{}
	//初始化缓存
	init_cache()
	//初始化文件变动监听
	init_watcher()
	// 遍历主目录生成菜单数据
	GenerateMenu()
	//初始化搜索引擎
	init_serach_engine()

	//执行创建索引
	indexing()
}

func init_cache() {
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 2048,
		// time after which entry can be evicted
		LifeWindow:       60 * 24 * time.Minute,
		HardMaxCacheSize: 1024,
	}

	ci, initErr := bigcache.NewBigCache(config)
	if initErr != nil {
		log.Fatal(initErr)
	}
	Cache = ci
}

func file_change_handler() {
	for {
		select {
		case event := <-watcher.Events:
			// 粗暴
			//目标变更重新生成菜单,并清空缓存
			GenerateMenu()
			Cache.Reset()
			log.Info(event)
		case err := <-watcher.Errors:
			log.Error("error:", err)
		}
	}

}

// 初始化监听器，监控文件变化
func init_watcher() {

	w, err := fsnotify.NewWatcher()
	watcher = w

	if err != nil {
		log.Fatal(err)
	}

	//defer watcher.Close()
	go file_change_handler()

	err = watcher.Add(Config.HomeDir)
	if err != nil {
		log.Fatal(err)
	}

}

// 初始化索引
func init_serach_engine() {

	mapping := bleve.NewIndexMapping()
	var err error

	DocIndex, err = bleve.Open(Config.IndexDir)

	if err != nil {
		log.Info(err)
		DocIndex, err = bleve.New(Config.IndexDir, mapping)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Info("create new index file in " + Config.IndexDir)
		}
	}
	log.Info(DocIndex.StatsMap())

}
