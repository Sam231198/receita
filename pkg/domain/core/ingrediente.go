package core

type Ingrediente struct {
	Id        int32   `json:"id,omitempty"`
	Nome      string  `json:"nome"`
	Qtd       float32 `json:"qtd,omitempty"`
	UndMedida string  `json:"undMedida"`
}

func NewIngrediente(nome string, undMedida string) Ingrediente {
	return Ingrediente{
		Nome:      nome,
		UndMedida: undMedida,
	}
}
