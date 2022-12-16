package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
)

type Student struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Level      string `json:"level"`
}

var students = []Student{
	{
		ID:         "10000xbcd3",
		Name:       "Alicia Winds",
		Department: "Political Science",
		Level:      "Year 3",
	},
}

func main() {
	router := gin.Default()
	router.GET("/", WelcomeMessage)
	router.POST("/createStudent", CreateStudent())
	router.GET("/students", GetStudents())
	router.Run()
}

func WelcomeMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hey boss!"})
}

func CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newStudent Student
		if err := c.BindJSON(&newStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}
		newStudent.ID = xid.New().String()
		students = append(students, newStudent)
		c.JSON(http.StatusCreated, newStudent)
	}
}

func GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Fetch all students in the DB
		c.JSON(http.StatusOK, students)
	}
}
