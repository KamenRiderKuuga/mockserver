//go:build release

package main

import "github.com/gin-gonic/gin"

func init() {
    // release模式配置
    gin.SetMode(gin.ReleaseMode)
}