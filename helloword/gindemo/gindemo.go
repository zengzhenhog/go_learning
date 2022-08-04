package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	const keyRequestId = "requestId"
	// middleware
	r.Use(func(ctx *gin.Context) {
		s := time.Now()
		ctx.Next()

		logger.Info("incoming request",
			zap.String("path", ctx.Request.URL.Path),
			zap.Int("status", ctx.Writer.Status()),
			// zap.Duration("elapsed", time.Now().Sub(s)))
			zap.Duration("elapsed", time.Since(s)))
	}, func(ctx *gin.Context) {
		// ctx.Set("requestId", rand.Int())
		ctx.Set(keyRequestId, 4324365656)
		ctx.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}

		c.JSON(http.StatusOK, h)
	})
	r.GET("hello", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})
	r.Run()
}
