package main

import (
	"github.com/gin-gonic/gin"
	"pruebas/configs"
	"pruebas/controllers"
	"pruebas/migrate"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
	migrate.SyncDataBase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes := r.Group("/api/v1/persons")
	routes.POST("", controllers.CreatePerson)
	routes.GET("", controllers.GetPersons)
	routes.GET("/:id", controllers.GetPersonById)
	routes.PUT("/:id", controllers.UpdatePerson)
	routes.DELETE("/:id", controllers.DeletePerson)
	r.Run() // listen and serve on 0.0.0.0:8080
}
