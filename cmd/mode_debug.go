//go:build !release

package main

import "github.com/gin-gonic/gin"

func init() {
    // debug模式配置
    gin.SetMode(gin.DebugMode)
}