package repositories

import (
	"database/sql"

	"github.com/samuel/receita/pkg/domain/core"
)

type IngredienteRepository struct {
	db *sql.DB
}

func NewIngredienteRepository(db *sql.DB) IngredienteRepository {
	return IngredienteRepository{db}
}

func (ir IngredienteRepository) FetchOne(ingredienteId int32) (core.Ingrediente, error) {
	row := ir.db.QueryRow(`SELECT * FROM ingredientes WHERE id = $1`, ingredienteId)

	var id int32
	var nome string
	var undMedida string

	err := row.Scan(
		&id,
		&nome,
		&undMedida,
	)

	if err != nil {
		return core.Ingrediente{}, err
	}

	ingrediente := core.NewIngrediente(nome, undMedida)
	ingrediente.Id = id

	return ingrediente, nil
}

// func (ir IngredienteRepository) FetchReceita(idReceita int) ([]core.Ingrediente, error) {
// }

func (ir IngredienteRepository) FetchAll() ([]core.Ingrediente, error) {
	rows, err := ir.db.Query("SELECT * FROM ingredientes")

	if err != nil {
		return []core.Ingrediente{}, err
	}

	ingredientes := []core.Ingrediente{}
	for rows.Next() {
		var id int32
		var nome string
		var undMedida string

		rows.Scan(
			&id,
			&nome,
			&undMedida,
		)

		ingrediente := core.NewIngrediente(nome, undMedida)
		ingrediente.Id = id

		ingredientes = append(ingredientes, ingrediente)
	}

	return ingredientes, nil
}

func (ir IngredienteRepository) Insert(ingrediente core.Ingrediente) (int32, error) {
	row := ir.db.QueryRow("INSERT INTO ingredientes (nome, und_medida) VALUES ($1,$2) RETURNING id",
		ingrediente.Nome, ingrediente.UndMedida)

	var id int32
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ir IngredienteRepository) Update(ingrediente core.Ingrediente) (int32, error) {
	row := ir.db.QueryRow("UPDATE ingredientes SET nome = $1, und_medida = $2 WHERE id = $3 RETURNING id",
		ingrediente.Nome, ingrediente.UndMedida, ingrediente.Id)

	var id int32
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
