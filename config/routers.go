package config

import (
	"github.com/gin-gonic/gin"
	"github.com/samuel/receita/pkg/controllers"
	"github.com/samuel/receita/pkg/repositories"
)

func LoadControllers(server *gin.Engine) {
	db := ConnectDB()

	// repositories
	receitaRepository := repositories.NewReceitaRepository(db)
	// controllers
	receitaController := controllers.NewReceitaController(receitaRepository)

	server.GET("/receita/lista", receitaController.ListReceita)
	server.POST("/receita/add", receitaController.AddReceita)
}
