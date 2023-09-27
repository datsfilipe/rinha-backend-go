package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/datsfilipe/rinha-backend-go/pkg/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/pessoas", func(c *gin.Context) {
		contextBody, _ := c.GetRawData()
		response, err := handlers.CreatePersonHandler(contextBody)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.String(http.StatusOK, string(response))
	})

	return r
}
