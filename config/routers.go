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
	ingredienteRepository := repositories.NewIngredienteRepository(db)
	// controllers
	receitaController := controllers.NewReceitaController(receitaRepository)
	ingredienteController := controllers.NewIngredienteController(ingredienteRepository)

	server.GET("/receita/:receitaID", receitaController.SearchReceita)
	server.GET("/receita/lista", receitaController.ListReceita)
	server.POST("/receita/add", receitaController.AddReceita)

	server.GET("/ingrediente/:receitaID", ingredienteController.SearchIngrediente)
	server.GET("/ingrediente/lista", ingredienteController.ListIngrediente)
	server.POST("/ingrediente/add", ingredienteController.AddIngrediente)
}
