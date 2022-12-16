package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestWelcomeMessage(t *testing.T) {
	mockResponse := `{"message":"Hey boss!"}`
	r := SetRouter()
	r.GET("/", WelcomeMessage)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateStudent(t *testing.T) {
	r := SetRouter()
	r.POST("/createStudent", CreateStudent())
	studentId := xid.New().String()
	student := Student{
		ID:         studentId,
		Name:       "Greg Winds",
		Department: "Political Science",
		Level:      "Year 4"}
	jsonValue, _ := json.Marshal(student)
	req, _ := http.NewRequest("POST", "/createStudent", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetStudents(t *testing.T) {
	r := SetRouter()
	r.GET("/students", GetStudents())
	req, _ := http.NewRequest("GET", "/students", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var students []Student
	json.Unmarshal(w.Body.Bytes(), &students)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, students)
}
