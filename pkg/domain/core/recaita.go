package core

type Receita struct {
	Id          int32         `json:"id,omitempty"`
	Nome        string        `json:"nome"`
	Tempo       int32         `json:"tempo,omitempty"`
	Rendimento  string        `json:"rendimento,omitempty"`
	Ingrediente []Ingrediente `json:"ingrediente,omitempty"`
}

func NewReceita(nome string) Receita {
	return Receita{
		Nome: nome,
	}
}