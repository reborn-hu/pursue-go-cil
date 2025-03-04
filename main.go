package main

import (
	"github.com/reborn-hu/gsharp"
	"os"
)

// @title  API接口文档
// @version 1.0
// @description  API接口文档
// @host localhost:50001
// @BasePath /
func main() {
	var builder = gsharp.CreateHostBuilder(&gsharp.HostOptions{EnvironmentName: os.Getenv("APP_ENV")})

	builder.UseConfiguration("./settings/application.json")
	builder.UseLogger()
	builder.UseDatabaseBuilder()
	builder.UseRedisBuilder()
	builder.UseLocalCacheBuilder()

	var app = builder.WebBuild()

	//app.UseRouter("/", modules.GetRouter())

	app.Static("/assets", "./static/assets")
	app.StaticFile("/", "./static/index.html")
	app.StaticFile("/favicon.ico", "./static/favicon.ico")

	app.Run()
}
