package main

import (
	"ginproject/routers"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof start fail")
			panic(err)
		}
	}()
	
	r := gin.Default()

	//加载模版
	r.LoadHTMLGlob("./templates/**/*.html")

	//加载静态文件
	r.Static("/statics", "./statics")

	//初始化路由
	//初始化默认界面
	routers.DefaultRountersInit(r)

	r.Run(":80")
	
}
