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
	server.GET("/receita", receitaController.ListReceita)
	server.POST("/receita", receitaController.AddReceita)
	server.PUT("/receita/:receitaID", receitaController.UpdReceita)

	server.GET("/ingrediente/:ingredienteID", ingredienteController.SearchIngrediente)
	server.GET("/ingrediente", ingredienteController.ListIngrediente)
	server.POST("/ingrediente", ingredienteController.AddIngrediente)
	server.PUT("/ingrediente/:ingredienteID", ingredienteController.UpdIngrediente)
}
