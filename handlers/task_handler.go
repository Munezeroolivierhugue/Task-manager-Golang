package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Munezeroolivierhugue/Task-manager-Golang/database"
	"github.com/Munezeroolivierhugue/Task-manager-Golang/models"
)

func PingHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message": "pong",})
}

func GetTask(c *gin.Context){
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK,tasks)
}

func GetTaskById(c *gin.Context){
	idParam := c.Param("id")
	id,err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var tasks models.Task
	result := database.DB.First(&tasks,id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound,gin.H{"Error": "task not found"})
			return
		}
	c.JSON(http.StatusOK,tasks)
}

func UpdateTask(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error": "Invalid id"})
		return
	}

	var task models.Task
	if result := database.DB.First(&task,id); result.Error != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	} 

	var updatedData models.Task
	if err := c.ShouldBindJSON(&updatedData); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Title = updatedData.Title
	task.Completed = updatedData.Completed
	database.DB.Save(&task)

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var task models.Task
	if result := database.DB.First(&task,id); result.Error != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

func CreateTask(c *gin.Context){
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newTask)
	c.JSON(http.StatusCreated, newTask)
}