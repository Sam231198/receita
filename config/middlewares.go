package config

import (
	"github.com/gin-gonic/gin"
	"github.com/samuel/receita/pkg/middlewares"
)

func LoadMiddlewares(engine *gin.Engine) {
	engine.Use(middlewares.CorsMiddleware())
}
