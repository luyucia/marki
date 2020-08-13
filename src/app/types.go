package app

import (
	"github.com/allegro/bigcache"
	"github.com/blevesearch/bleve"
	"github.com/fsnotify/fsnotify"
	"sync"
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
var GConfig = AppConfig{}

//全局变量
var GMenuData FileNode
var GDocumentMap map[string]Document
var GCache *bigcache.BigCache
var GWatcher *fsnotify.Watcher
var DocIndex bleve.Index

// 上次菜单更新时间
var GLastMenuUpdateTime = time.Unix(10000, 1)
var GMenuGenerateMutex = sync.Mutex{}
