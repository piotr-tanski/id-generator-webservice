package main

import (
	"id_generator/generator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getNextUniqueIdHandler(generator *generator.IdGenerator) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, timestamp, machine_id, seqno := generator.Generate()
		c.IndentedJSON(http.StatusOK, gin.H{"id": id, "timestamp": timestamp, "machine_id": machine_id, "seqno": seqno})
	}
	return gin.HandlerFunc(fn)
}

func main() {
	machine_id := 1
	generator := generator.NewGenerator(machine_id)

	router := gin.Default()
	router.GET("/v1/ids", getNextUniqueIdHandler(&generator))
	router.Run("localhost:8080")
}
