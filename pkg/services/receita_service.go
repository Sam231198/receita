package services

import (
	"fmt"

	"github.com/samuel/receita/pkg/domain/core"
	"github.com/samuel/receita/pkg/domain/repository"
)

type ReceitaService struct {
	receitaRepository repository.ReceitaRepository
}

func NewReceitaService(receitaRepositor repository.ReceitaRepository) ReceitaService {
	return ReceitaService{receitaRepositor}
}

func (rs ReceitaService) GetAll() []core.Receita {
	receitas, err := rs.receitaRepository.FetchAll()
	if err != nil {
		fmt.Print("Error ao retornar as receitas do banco de dados: ", err.Error())
		return []core.Receita{}

	}
	return receitas
}

func (rs ReceitaService) Create(receita core.Receita) core.Receita {
	id, err := rs.receitaRepository.Insert(receita)
	if err != nil {
		fmt.Print("A inserção falhou: ", err.Error())
		return core.Receita{}
	}

	receita.Id = id
	return receita

}
