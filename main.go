package main

import (
	"goLog/logs"
	"goapiv1/config"
	"goapiv1/model"

	"github.com/labstack/echo"
)

var (
	e = echo.New()
)

func main() {
	RegisterApi()
	config.InitConfig()
	logs.Debug.Println(config.Config.DB.Conn)
	model.InitDB("mysql", config.Config.DB.Conn)
	e.Start(":1114")
}
