package main

import (
	"flag"
	"fmt"
	"log"
	"mockserver/internal/config"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// if not exists, create a new directory named "static"
	if _, err := os.Stat("static"); os.IsNotExist(err) {
		os.Mkdir("static", 0755)
	}
	// Load configuration
	routesConfig, err := config.LoadConfig("static")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Gin engine
	r := gin.Default()

	// Register routes
	config.RegisterRoutes(r, routesConfig)

	// get the port from command line
	// 从命令行读取端口参数
	port := flag.String("port", "80", "端口号")
	flag.Parse()

	// 使用命令行参数中的端口启动
	r.Run(fmt.Sprintf(":%s", *port))
}
