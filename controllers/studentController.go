package controllers

import (
	"bibi/config"
	"bibi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @summary  get all students
// @tags students
// @router /student/ [get]
func GetAllStudent(c *gin.Context) {
	var products []models.Students

	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// @summary create student
// @tags students
// @param student body models.Students true "student data"
// @router /student/ [post]
func CreateStudent(c *gin.Context) {
	var input models.Students
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Students{
		Name:    input.Name,
		Age:     input.Age,
		Class:   input.Class,
		Address: input.Address,
	}

	config.DB.Create(&student)
	c.JSON(http.StatusCreated, gin.H{"student": student})
}

// @summary update student
// @tags students
// @param student_id path int true "StudentID"
// @param student body models.Students true "Student info"
// @router /student/{student_id} [put]
func UpdateStudent(c *gin.Context) {
	var input models.Students

	// Gán giá trị JSON vào input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lấy thông tin sinh viên hiện tại bằng id
	var student models.Students
	if err := config.DB.Where("student_id = ?", c.Param("student_id")).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Cập nhật thông tin sinh viên
	student.Name = input.Name
	student.Age = input.Age
	student.Class = input.Class
	student.Address = input.Address

	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"student updated": student})
}

// @summary delete student
// @tags students
// @param student_id path int true "StudentID"
// @router /student/{student_id} [delete]
func DeleteStudent(c *gin.Context) {
	var student models.Students
	if err := config.DB.Where("student_id = ?", c.Param("student_id")).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	config.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"data": "da xoa thanh cong"})
}

// @summary get student by id
// @tags students
// @param student_id path int true "StudentID"
// @router /student/{student_id} [get]
func GetStudentById(c *gin.Context) {
	var student models.Students
	if err := config.DB.Where("student_id = ?", c.Param("student_id")).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"student": student})
}
