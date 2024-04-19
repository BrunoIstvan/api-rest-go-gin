package routes

import (
	"github.com/BrunoIstvan/api-rest-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/students", controllers.ListAllStudents)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)
	r.GET("/students/:id", controllers.GetStudentById)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudentById)
	r.PATCH("/students/:id", controllers.UpdateStudentById)
	r.GET("/index", controllers.LoadIndexPage)
	r.NoRoute(controllers.LoadNotFoundPage)

	r.Run()
}
