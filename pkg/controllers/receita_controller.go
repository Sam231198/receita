package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuel/receita/pkg/domain/core"
	"github.com/samuel/receita/pkg/domain/repository"
	"github.com/samuel/receita/pkg/services"
)

type ReceitaController struct {
	sv services.ReceitaService
}

func NewReceitaController(receitaRepository repository.ReceitaRepository) ReceitaController {
	return ReceitaController{sv: services.NewReceitaService(receitaRepository)}
}

func (rc ReceitaController) ListReceita(ctx *gin.Context) {
	receitas := rc.sv.GetAll()
	ctx.JSON(http.StatusOK, gin.H{"result": receitas})
}

func (rc ReceitaController) AddReceita(ctx *gin.Context) {
	receita := core.Receita{}
	err := ctx.ShouldBindJSON(&receita)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível ler os dados da receita"})
		return
	}

	receita = rc.sv.Create(receita)
	
	ctx.JSON(http.StatusCreated, gin.H{"result": receita})
}
