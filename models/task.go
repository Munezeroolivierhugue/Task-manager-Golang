package models

type Task struct{
	ID uint `gorm:"primarykey" json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}