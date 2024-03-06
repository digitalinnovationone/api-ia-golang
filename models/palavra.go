package models

type Palavra struct {
	Palavra  string   `json:"palavra"`
	Traducao string   `json:"traducao"`
	Opcoes   []string `json:"opcoes"`
}
