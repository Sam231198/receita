package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuel/receita/pkg/domain/core"
	"github.com/samuel/receita/pkg/domain/repository"
	"github.com/samuel/receita/pkg/services"
)

type IngredienteController struct {
	si services.IngredienteService
}

func NewIngredienteController(ingredienteRepository repository.IngredienteRepository) IngredienteController {
	return IngredienteController{services.NewIngredienteService(ingredienteRepository)}
}

func (ic IngredienteController) SearchIngrediente(ctx *gin.Context) {
	ingredienteId, err := strconv.Atoi(ctx.Param("ingredienteID"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"receitaID": "O valor do parametro não esta correto, deve ser um número inteiro"})
	}

	ingrediente := ic.si.GetOne(int32(ingredienteId))
	ctx.JSON(http.StatusOK, gin.H{"ingrediente": ingrediente})
}

func (ic IngredienteController) ListIngrediente(ctx *gin.Context) {
	ingredientes := ic.si.GetAll()
	ctx.JSON(http.StatusOK, gin.H{"ingredientes": ingredientes})
}

func (ic IngredienteController) AddIngrediente(ctx *gin.Context) {
	ingrediente := core.Ingrediente{}
	err := ctx.ShouldBindJSON(&ingrediente)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível ler os dados do ingrediente"})
		return
	}

	ingrediente = ic.si.Create(ingrediente)
	ctx.JSON(http.StatusCreated, gin.H{"ingrediente": ingrediente})
}
