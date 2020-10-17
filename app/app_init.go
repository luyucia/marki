package app

import (
	"flag"
	"github.com/allegro/bigcache"
	"github.com/blevesearch/bleve"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/gommon/log"
	"os"
	"time"
)

// init 按照文件名顺序执行，为了让这个先执行，所以用首字母a,叫app_init

//初始化
func init() {
	// 初始化默认参数
	GConfig.Host = "0.0.0.0"
	GConfig.Port = 80
	//GConfig.HomeDir = "h:\\mydocs"
	GConfig.HomeDir = "./"
	GConfig.CacheDir = "./cache"
	GConfig.IndexDir = "./index.bleve"
	GConfig.AllowType = map[string]bool{}
	GConfig.AllowType["md"] = true
	GConfig.AllowType["http"] = true
	GConfig.AllowType["map"] = true

	flag.StringVar(&GConfig.Host, "host", "0.0.0.0", "监听地址")
	flag.IntVar(&GConfig.Port, "port", 80, "端口")
	flag.StringVar(&GConfig.HomeDir, "path", "./", "文档地址")
	flag.Parse()

	if len(os.Args) > 1 {
		GConfig.HomeDir = os.Args[1]
	}

	log.Info("Home path:", GConfig.HomeDir)
	log.Info("Listen port:", GConfig.Port)
	log.Info("Listen host:", GConfig.Host)
	log.Info("Allow Type:", GConfig.AllowType)

	// 初始化目录结构
	GDocumentMap = map[string]Document{}
	//初始化缓存
	init_cache()
	//初始化文件变动监听
	init_watcher()
	// 遍历主目录生成菜单数据
	GenerateMenu()
	//初始化搜索引擎
	//init_serach_engine()
	//执行创建索引
	//indexing()
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
	GCache = ci
}

func file_change_handler() {
	for {
		select {
		case event := <-GWatcher.Events:
			// 粗暴
			//目标变更重新生成菜单,并清空缓存
			GenerateMenu()
			GCache.Reset()
			log.Info(event)
		case err := <-GWatcher.Errors:
			log.Error("error:", err)
		}
	}

}

// 初始化监听器，监控文件变化
func init_watcher() {

	w, err := fsnotify.NewWatcher()
	GWatcher = w

	if err != nil {
		log.Fatal(err)
	}

	//defer GWatcher.Close()
	go file_change_handler()

	err = GWatcher.Add(GConfig.HomeDir)
	if err != nil {
		log.Fatal(err)
	}

}

// 初始化索引
func init_serach_engine() {

	mapping := bleve.NewIndexMapping()
	var err error

	DocIndex, err = bleve.Open(GConfig.IndexDir)

	if err != nil {
		log.Info(err)
		DocIndex, err = bleve.New(GConfig.IndexDir, mapping)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Info("create new index file in " + GConfig.IndexDir)
		}
	}
	log.Info(DocIndex.StatsMap())

}
