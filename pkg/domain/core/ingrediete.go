package core

type Ingrediente struct {
	Nome      string  `json:"nome"`
	Qtd       float32 `json:"qtd"`
	UndMedida string  `json:"undMedida"`
}

func NewIngrediente(nome string, qtd float32, undMedida string) *Ingrediente {
	return &Ingrediente{
		Nome:      nome,
		Qtd:       qtd,
		UndMedida: undMedida,
	}
}
