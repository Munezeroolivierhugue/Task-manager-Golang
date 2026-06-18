package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Task struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

var tasks = []Task{}
var nextID = 1

//ping handler/similar to controller
func PingHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetTask(c *gin.Context){
	c.JSON(http.StatusOK,tasks)
}

func GetTaskById(c *gin.Context){
	idParam := c.Param("id")
	id,err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	for _, t := range tasks{
		if t.ID == id {
			c.JSON(http.StatusOK,t)
			return
		}
	}
	c.JSON(http.StatusNotFound,gin.H{"Error": "task not found"})
}

func UpdateTask(c *gin.Context){
	idParam := c.Param("id")
	id,err := strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error": "Invalid id"})
		return
	}

	var updatedData Task
	if err := c.ShouldBindJSON(&updatedData); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range tasks{
		if tasks[i].ID == id{
			tasks[i].Title = updatedData.Title
			tasks[i].Completed = updatedData.Completed
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTask(c *gin.Context){
	idParam := c.Param("id")
	id,err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	for i := range tasks{
		if tasks[i].ID == id{
			tasks = append(tasks[:i],tasks[i+1:]...)
			c.JSON(http.StatusOK, tasks)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func CreateTask(c *gin.Context){
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newTask.ID = nextID
	nextID++
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}
/* 
pingHandler here accepts an argument of a reference to context of gin which is used for accepting a request in this case we have a response of JSON form, with http also imported in net/http used for codes too here
gin.H here is very crucial it makes an object simpler form of map[string]interface{}
*/

//main method is Go's entry point
func main_file_we_used_before(){
	r := gin.Default()
	r.GET("/ping", PingHandler)
	r.GET("/tasks", GetTask)
	r.GET("/tasks/:id", GetTaskById)
	r.POST("/tasks", CreateTask)
	r.PUT("/tasks/:id",UpdateTask)
	r.DELETE("/tasks/:id",DeleteTask)
	r.Run(":8080")
}