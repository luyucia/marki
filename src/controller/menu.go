package controller

import (
	"github.com/labstack/echo/v4"
	"marki/app"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(data interface{}) Response {
	r := Response{}
	r.Msg = "ok"
	r.Code = 0
	r.Data = data
	return r
}

func GetMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, SuccessResponse(app.MenuData))
}

func GetContent(c echo.Context) error {

	id := c.QueryParam("id")
	html := app.GenerateContent(id)
	return c.HTML(http.StatusOK,string(html))

}
