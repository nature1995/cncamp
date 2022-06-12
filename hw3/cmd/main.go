package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"hw1/cmd/httpserver"
	"hw1/internal/config"
	"hw1/internal/logger"
)

func main() {
	// 初始化配置
	err := config.GetConfig("")
	if err != nil {
		fmt.Printf("Load config fail: %s", err)
		os.Exit(1)
	}

	// 初始化日志
	log, err := logger.NewLogger(config.C.Log)
	if err != nil {
		fmt.Printf("Failed to create logger: %s", err)
		os.Exit(1)
	}

	// 初始化服务
	server := httpserver.NewHttpServer(config.C.HTTP)
	go func() {
		if err = server.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	// 优雅重启
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Info("Gracefully terminating...")

	timeout := time.Second * time.Duration(30)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Errorf("Shutting down err: %s", err)
	}

	log.Info("Gracefully Shutting down successfully")
}
