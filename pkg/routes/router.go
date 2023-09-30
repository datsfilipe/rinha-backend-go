package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/datsfilipe/rinha-backend-go/pkg/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/pessoas", func(c *gin.Context) {
		contextBody, _ := c.GetRawData()
		response, err := handlers.CreatePersonHandler(contextBody)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.String(http.StatusOK, string(response))
	})

	r.GET("/pessoas/:id", func(c *gin.Context) {
		id := c.Param("id")
		response, err := handlers.GetPersonHandler(id)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.String(http.StatusOK, string(response))
	})

	r.GET("/contagem-pessoas", func(c *gin.Context) {
		response, err := handlers.CountPeopleHandler()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.String(http.StatusOK, string(response))
	})

	return r
}
