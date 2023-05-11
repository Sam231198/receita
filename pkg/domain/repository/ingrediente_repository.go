package repository

import "github.com/samuel/receita/pkg/domain/core"

type IngredienteRepository interface {
	FetchOne(ingredienteId int32) (core.Ingrediente, error)
	// FetchReceita(idReceita int32) ([]core.Ingrediente, error)
	FetchAll() ([]core.Ingrediente, error)
	Insert(core.Ingrediente) (int32, error)
}
