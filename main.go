package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"we-talk-backend/initial"
)

var service *initial.Service

func main() {
	var err error
	service, err = initial.Initial()
	if err != nil {
		return
	}
	r := gin.Default()
	r.GET("/register", handleRegister)
	err = r.Run()
	if err != nil {
		return
	}
}

func handleRegister(c *gin.Context) {
	nickname := c.Query("nickname")
	password := c.Query("password")
	userId, status := service.AccountService.Register(nickname, password)
	c.JSON(
		200, gin.H{
			"userId":               userId,
			"registerReturnStatus": status,
		})
}
