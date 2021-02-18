package main

import (
	"github.com/PaleBlueYk/smileHugo-backend/config"
	"github.com/PaleBlueYk/smileHugo-backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")
	if err := config.Init(); err != nil {
		logger.Error(err)
		return
	}
}
