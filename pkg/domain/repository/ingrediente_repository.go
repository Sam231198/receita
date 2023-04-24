package repository

import "github.com/samuel/receita/pkg/domain/core"

type IngredienteRepository interface {
	FetchOne(id int) (core.Ingrediente, error)
	FetchReceita(idReceita int) ([]core.Ingrediente, error)
	FetchAll() ([]core.Ingrediente, error)
	Insert(core.Ingrediente) (core.Ingrediente, error)
}
