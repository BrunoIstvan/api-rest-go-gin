package controllers

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/BrunoIstvan/api-rest-go-gin/dtos"
	"github.com/BrunoIstvan/api-rest-go-gin/services"
	"github.com/gin-gonic/gin"
)

func ListAllStudents(c *gin.Context) {

	students := services.ListAllStudents()
	c.JSON(http.StatusOK, students)

}

func GetStudentById(c *gin.Context) {

	strId := c.Params.ByName("id")
	id := parseAndValidateId(c, strId, `^[0-9]+$`)
	if id == 0 {
		return
	}
	response := getStudentByIdAndValidateResponse(c, id)
	if response == nil {
		return
	}
	c.JSON(http.StatusOK, response)

}

func SearchStudentByCPF(c *gin.Context) {

	cpf := c.Params.ByName("cpf")
	if !validateIfItIsOnlyNumber(c, cpf, "cpf", `^[0-9]{11}$`) {
		return
	}
	response, err := services.SearchStudentByCPF(cpf)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}

func CreateStudent(c *gin.Context) {

	var request dtos.StudentRequestDTO
	if !bindStudentData(c, &request) {
		return
	}
	response, err := services.CreateStudent(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)

}

func DeleteStudentById(c *gin.Context) {

	strId := c.Params.ByName("id")
	id := parseAndValidateId(c, strId, `^[0-9]+$`)
	if id == 0 {
		return
	}
	err := services.DeleteStudentById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": "Student deleted successfuly!"})

}

func UpdateStudentById(c *gin.Context) {

	strId := c.Params.ByName("id")
	id := parseAndValidateId(c, strId, `^[0-9]+$`)
	if id == 0 {
		return
	}
	var request dtos.StudentRequestDTO
	if !bindStudentData(c, &request) {
		return
	}
	response, err := services.UpdateStudentById(id, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}

func LoadIndexPage(c *gin.Context) {

	students := services.ListAllStudents()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})

}

func LoadNotFoundPage(c *gin.Context) {

	c.HTML(http.StatusNotFound, "404.html", nil)

}

func validateIfItIsOnlyNumber(c *gin.Context, value string, fieldName string, regExp string) bool {
	var re = regexp.MustCompile(regExp)
	if !re.MatchString(value) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fieldName + "is invalid"})
		return false
	}
	return true
}

func getStudentByIdAndValidateResponse(c *gin.Context, id uint64) *dtos.StudentResponseDTO {
	response, err := services.GetStudentById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return nil
	}
	return response
}

func parseAndValidateId(c *gin.Context, strId string, regExp string) uint64 {
	var re = regexp.MustCompile(regExp)
	if !re.MatchString(strId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is invalid"})
		return 0
	}
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	return id
}

func bindStudentData(c *gin.Context, student *dtos.StudentRequestDTO) bool {

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true

}
