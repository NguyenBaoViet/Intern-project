package main

import (
	"Intern-project/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := gin.Default()
	r.GET("/user/:name", controller.Getuser)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:name", controller.UpdateUser)
	r.DELETE("/user/:name", controller.DeleteUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
