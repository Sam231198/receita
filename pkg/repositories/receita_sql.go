package repositories

import (
	"database/sql"

	"github.com/samuel/receita/pkg/domain/core"
)

type ReceitaRepository struct {
	db *sql.DB
}

func NewReceitaRepository(db *sql.DB) ReceitaRepository {
	return ReceitaRepository{db}
}

func (rp ReceitaRepository) FetchOne(receitaID int32) (core.Receita, error) {
	row := rp.db.QueryRow(`SELECT * FROM receitas WHERE id = $1`, receitaID)

	var id int32
	var nome string
	var tempo int32
	var rendimento string

	err := row.Scan(
			&id,
			&nome,
			&tempo,
			&rendimento,
		)

	if err != nil {
		return core.Receita{}, err
	}

	receita := core.NewReceita(nome)
	receita.Id = id
	receita.Tempo = tempo
	receita.Rendimento = rendimento

	return receita, nil
}

func (rp ReceitaRepository) FetchAll() ([]core.Receita, error) {
	rows, err := rp.db.Query("SELECT * FROM receitas")
	if err != nil {
		return []core.Receita{}, err
	}
	defer rows.Close()

	receitas := []core.Receita{}
	for rows.Next() {
		var id int32
		var nome string
		var tempo int32
		var rendimento string

		rows.Scan(
			&id,
			&nome,
			&tempo,
			&rendimento,
		)

		receita := core.NewReceita(nome)
		receita.Id = id
		receita.Tempo = tempo
		receita.Rendimento = rendimento

		receitas = append(receitas, receita)
	}

	return receitas, nil
}

func (rp ReceitaRepository) Insert(receita core.Receita) (int32, error) {
	row := rp.db.QueryRow(`INSERT INTO receitas (nome, tempo, rendimento) VALUES ($1,$2,$3) RETURNING id`,
		receita.Nome.rstrip(' ').lstrip(' '), receita.Tempo, receita.Rendimento) 

	var id int32
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
