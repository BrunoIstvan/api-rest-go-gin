package services

import (
	"errors"
	"fmt"

	"github.com/BrunoIstvan/api-rest-go-gin/database"
	"github.com/BrunoIstvan/api-rest-go-gin/dtos"
	"github.com/BrunoIstvan/api-rest-go-gin/models"
)

func ListAllStudents() []dtos.StudentResponseDTO {
	var students []models.Student
	database.DB.Order("name ASC").Find(&students)
	var studentsDTOs []dtos.StudentResponseDTO
	for _, studentModel := range students {
		student := convertStudentModelToDTO(&studentModel)
		studentsDTOs = append(studentsDTOs, *student)
	}
	return studentsDTOs
}

func GetStudentById(id uint64) (*dtos.StudentResponseDTO, error) {

	var student models.Student
	database.DB.First(&student, id)
	if err := validateStudentExists(&student); err != nil {
		return nil, err
	}
	return convertStudentModelToDTO(&student), nil

}

func SearchStudentByCPF(cpf string) (*dtos.StudentResponseDTO, error) {

	var student models.Student
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)
	if err := validateStudentExists(&student); err != nil {
		return nil, err
	}
	return convertStudentModelToDTO(&student), nil

}

func CreateStudent(studentDTO *dtos.StudentRequestDTO) (*dtos.StudentResponseDTO, error) {

	if err := dtos.ValidateStudentData(studentDTO); err != nil {
		return nil, err
	}
	studentModel := convertStudentDTOToModel(studentDTO)
	database.DB.Create(&studentModel)
	return convertStudentModelToDTO(studentModel), nil

}

func DeleteStudentById(id uint64) error {
	if _, err := GetStudentById(id); err != nil {
		return err
	}
	var studentModel models.Student
	if result := database.DB.Delete(&studentModel, id); result == nil {
		return errors.New("Error while deleting Student!")
	}
	return nil

}

func UpdateStudentById(id uint64, studentDTO *dtos.StudentRequestDTO) (*dtos.StudentResponseDTO, error) {
	if err := dtos.ValidateStudentData(studentDTO); err != nil {
		fmt.Print("Saindo por conta de erro na validação dos dados de entrada!", err)
		return nil, err
	}
	if _, err := GetStudentById(id); err != nil {
		return nil, err
	}
	studentModel := convertStudentDTOToModel(studentDTO)
	studentModel.ID = uint(id)
	if result := database.DB.Save(studentModel); result == nil {
		return nil, errors.New("Error while updating Student!")
	}
	dto := convertStudentModelToDTO(studentModel)
	return dto, nil
}

func validateStudentExists(student *models.Student) error {
	if student.ID == 0 {
		return errors.New("Student not found!")
	}
	return nil
}

func convertStudentModelToDTO(model *models.Student) *dtos.StudentResponseDTO {
	return &dtos.StudentResponseDTO{
		ID:   model.ID,
		Name: model.Name,
		CPF:  model.CPF,
		RG:   model.RG,
	}
}

func convertStudentDTOToModel(dto *dtos.StudentRequestDTO) *models.Student {
	return &models.Student{
		Name: dto.Name,
		CPF:  dto.CPF,
		RG:   dto.RG,
	}
}
