package main

import (
	"id_generator/generator"
	"id_generator/handlers"

	"github.com/gin-gonic/gin"
)

func getNextUniqueIdHandler(handler *handlers.GenerateUniqueIdHandler) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		status, response := handler.GenerateId()
		c.IndentedJSON(status, response)
	}
	return gin.HandlerFunc(fn)
}

func main() {
	machine_id := 1
	generator := generator.NewGenerator(machine_id)
	uniqueIdHandler := handlers.GenerateUniqueIdHandler{Generator: &generator}

	router := gin.Default()
	router.GET("/v1/ids", getNextUniqueIdHandler(&uniqueIdHandler))
	router.Run("localhost:8081")
}
