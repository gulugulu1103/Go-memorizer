package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gulugulu1103/Go-memorizer/handler"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("account service started...")
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// 创建一个默认的路由引擎，使用了 Logger、Recovery 中间件
	router := gin.Default()

	// 路由 GET 请求，第一个参数是路径，第二个参数是处理这个请求的函数
	// 函数的要求是 func(c *gin.Context)，gin.Context 封装了 request 和 response
	// 这里返回一个 JSON，JSON 是一个 map[string]interface{} 的实例
	handler.NewHandler(&handler.Config{
		R: router,
	})

	// 为什么是&http.Server{}，而不是http.Server{}？
	//因为http.Server是一个结构体，而不是一个指针，所以不能直接传递给ListenAndServe方法，需要取地址
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
