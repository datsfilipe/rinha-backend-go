package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/datsfilipe/rinha-backend-go/pkg/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/pessoas", func(c *gin.Context) {
		contextBody, _ := c.GetRawData()
		response, status, err := handlers.CreatePersonHandler(contextBody)
		if err != nil {
			c.String(status, "")
		}
		c.String(status, string(response))
	})

	r.GET("/pessoas/:id", func(c *gin.Context) {
		id := c.Param("id")
		response, status, err := handlers.GetPersonHandler(id)
		if err != nil {
			c.String(status, "")
		}
		c.String(status, string(response))
	})

	r.GET("/pessoas", func(c *gin.Context) {
		t := c.Query("t")
		response, status, err := handlers.SearchByTermHandler(t)
		if err != nil {
			c.String(status, "")
		}
		c.String(status, string(response))
	})

	r.GET("/contagem-pessoas", func(c *gin.Context) {
		response, status, err := handlers.CountPeopleHandler()
		if err != nil {
			c.String(status, "")
		}
		c.String(status, string(response))
	})

	return r
}
