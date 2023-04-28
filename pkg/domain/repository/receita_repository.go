package repository

import "github.com/samuel/receita/pkg/domain/core"

type ReceitaRepository interface {
	FetchOne(id int32) (core.Receita, error)
	FetchAll() ([]core.Receita, error)
	Insert(core.Receita) (int32, error)
}
