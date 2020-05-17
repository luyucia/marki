package main

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tylerb/graceful"
	"marki/app"
	"marki/controller"
	"os"
	"strconv"
	"time"
)
import "github.com/labstack/echo/v4"


func init_config(e *echo.Echo) {

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetOutput(os.Stdout)
}

func main()  {
	e := echo.New()
	// 注册中间件
	RegisterMiddleware(e)
	// 注册路由
	RegisterRouter(e)
	// Start server
	init_config(e)

	e.Logger.Info("server start on http://" + app.Config.Host + ":" + strconv.Itoa(app.Config.Port))

	// 平滑关闭
	err := graceful.ListenAndServe(e.Server, 5*time.Second)
	if err != nil {
		e.Logger.Error(err)
	}
	e.Logger.Info("server stop success")
}


func RegisterRouter(e *echo.Echo) {

	//e.File("/", "web/dist/index.html")
	e.File("/", "web/index.html")
	e.Static("/css/*", "web/dist/css")
	e.Static("/js/*", "web/dist/js")

	// Routes
	//e.GET("/find/:key", controller.FileList)
	//e.GET("/get_text/:key", controller.TextContent)
	//e.POST("/put_text/:key", controller.PutTextContent)
	//e.POST("/file/upload", controller.FileUpload)
	//e.GET("/file/list", controller.FileList)
	//
	//e.GET("wsconncet", controller.WsConnect)
	e.GET("/get_menu",controller.GetMenu)
	e.GET("/get_content",controller.GetContent)
}

func RegisterMiddleware(e *echo.Echo) {
	// Middleware
	//e.Use(middleware.Logger()

	logfile, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		fmt.Printf("打开文件出错：%v\n", err)
	}

	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			//Format: "${time_unix},${level},${method},${uri},${status},${file},${message}\n",
			Output: logfile,
		}))
	e.Use(middleware.Recover())
}