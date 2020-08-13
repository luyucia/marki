package controller

import (
	"github.com/labstack/echo/v4"
	"marki/app"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Content struct {
	ContentText string `json:"content"`
	Type        string `json:"content_type"`
}

func SuccessResponse(data interface{}) Response {
	r := Response{}
	r.Msg = "ok"
	r.Code = 0
	r.Data = data
	return r
}

func GetMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, SuccessResponse(app.GMenuData))
}

func GetContent(c echo.Context) error {

	id := c.QueryParam("id")
	html, ext := app.GenerateContent(id)

	content := Content{}
	content.Type = ext
	content.ContentText = string(html)

	return c.JSON(http.StatusOK, content)

}
