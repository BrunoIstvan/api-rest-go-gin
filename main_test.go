package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoIstvan/api-rest-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTestsRoutes() *gin.Engine {

	routes := gin.Default()

	return routes

}

func TestWhenListAllStudentsShouldReturnStatusCode200(t *testing.T) {

	r := SetupTestsRoutes()
	r.GET("/students", controllers.ListAllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Fatalf("Status error: received value was %d, but expected was %d", response.Code, http.StatusOK)
	}

}
