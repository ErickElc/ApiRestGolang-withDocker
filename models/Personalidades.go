package models

type Personalidade struct {
	Id       int    `json: "id"`
	Nome     string `json: "nome"`
	Idade    int    `json: "idade"`
	Historia string `json: "historia"`
}

var Personalidades []Personalidade
