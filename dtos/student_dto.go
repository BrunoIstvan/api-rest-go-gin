package dtos

import "gopkg.in/validator.v2"

type StudentResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

type StudentRequestDTO struct {
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`
}

func ValidateStudentData(student *StudentRequestDTO) error {

	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil

}
