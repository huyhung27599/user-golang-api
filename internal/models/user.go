package models

type User struct {
	UUID string `json:"uuid" binding:"required,uuid"`	
	Name string `json:"name" binding:"required,min=3,max=100"`
	Age int `json:"age" binding:"required,gt=0,lt=100"`
	Email string `json:"email" binding:"required,email,email_address"`
	Password string `json:"password" binding:"required,min=8,max=20,password_strong"`
	Status int `json:"status" binding:"required,oneof=1 2"`
	Level int `json:"level" binding:"required,oneof=1 2 "`
}