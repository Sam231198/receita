package services

import (
	"fmt"

	"github.com/samuel/receita/pkg/domain/core"
	"github.com/samuel/receita/pkg/domain/repository"
)

type IngredienteService struct {
	ingredienteRepository repository.IngredienteRepository
}

func NewIngredienteService(ingredienteRepository repository.IngredienteRepository) IngredienteService {
	return IngredienteService{ingredienteRepository}
}

func (is IngredienteService) GetOne(ingredienteId int32) core.Ingrediente {
	ingrediente, err := is.ingredienteRepository.FetchOne(ingredienteId)

	if err != nil {
		fmt.Print("Error ao retornar o ingrediente do banco de dados: ", err.Error())
		return core.Ingrediente{}
	}

	return ingrediente
}

func (is IngredienteService) GetAll() []core.Ingrediente {
	ingredientes, err := is.ingredienteRepository.FetchAll()

	if err != nil {
		fmt.Print("Error ao retornar os ingredientes do banco de dados: ", err.Error())
		return []core.Ingrediente{}
	}

	return ingredientes
}

func (is IngredienteService) Create(ingrediente core.Ingrediente) core.Ingrediente {
	id, err := is.ingredienteRepository.Insert(ingrediente)

	if err != nil {
		fmt.Print("Error ao retornar os ingredientes do banco de dados: ", err.Error())
		return core.Ingrediente{}
	}

	ingrediente.Id = id
	return ingrediente
}

func (is IngredienteService) Update(ingrediente core.Ingrediente) core.Ingrediente {
	id, err := is.ingredienteRepository.Update(ingrediente)

	if err != nil {
		fmt.Print("Error ao retornar os ingredientes do banco de dados: ", err.Error())
		return core.Ingrediente{}
	}

	ingrediente.Id = id
	return ingrediente
}
