package main

import (
	"github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"

    "github.com/Munezeroolivierhugue/Task-manager-Golang/database"
	"github.com/Munezeroolivierhugue/Task-manager-Golang/handlers"
)

func main(){
	if err := godotenv.Load(); err != nil{
		log.Println("no .env file found, relying on system env vars")
	}

	database.Connect()

	r := gin.Default()

	r.GET("/ping", handlers.PingHandler)
    r.GET("/tasks", handlers.GetTask)
    r.GET("/tasks/:id", handlers.GetTaskById)
    r.POST("/tasks", handlers.CreateTask)
    r.PUT("/tasks/:id", handlers.UpdateTask)
    r.DELETE("/tasks/:id", handlers.DeleteTask)

	r.Run(":8080")
}