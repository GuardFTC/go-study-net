// Package gin @Author:冯铁城 [17615007230@163.com] 2023-06-25 14:38:54
package gin

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-study-net/gin/routers"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// StartGinServer gin框架开启服务器
func StartGinServer() {

	//1.初始化路由
	routers.InitRouter()

	//2.运行服务1
	g.Go(func() error {
		return routers.Router1.Run(":8080")
	})

	//3.运行服务2
	g.Go(func() error {
		return routers.Router2.Run(":8081")
	})

	//4.监听异常
	if err := g.Wait(); err != nil {
		log.Fatalf("server run err->[%v]\n", err)
	}
}

// 关闭服务器
func exitServer(server *http.Server) {

	//1.阻塞等待服务关闭
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	//2.服务关闭后等待5s退出
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	} else {
		log.Println("Server exiting")
	}
}
