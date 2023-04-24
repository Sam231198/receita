package repository

import "github.com/samuel/receita/pkg/domain/core"

type ReceitaRepository interface {
	FetchAll() ([]core.Receita, error)
	Insert(core.Receita) (int32, error)
	// FetchOne(id int) (core.Receita, error)
}
