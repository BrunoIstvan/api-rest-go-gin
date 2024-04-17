package models

type Student struct {
	MyModel
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Students []Student
