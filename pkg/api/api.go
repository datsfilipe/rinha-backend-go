package api

import (
	"database/sql"

	"github.com/datsfilipe/rinha-backend-go/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.POST("/pessoas", func(c *gin.Context) {
		contextBody, _ := c.GetRawData()
		response, status, err, location := handlers.CreatePersonHandler(db, contextBody)
		if err != nil {
			c.String(status, "")
		}
		c.Header("Location", location)
		c.String(status, string(response))
	})

	r.GET("/pessoas/:id", func(c *gin.Context) {
		id := c.Param("id")

		response, status, err := handlers.GetPersonHandler(db, id)
		if err != nil {
			c.String(status, "")
			return
		}
		c.String(status, string(response))
	})

	r.GET("/pessoas", func(c *gin.Context) {
		t := c.Query("t")
		response, status, err := handlers.SearchByTermHandler(db, t)
		if err != nil {
			c.String(status, "")
			return
		}
		c.String(status, string(response))
	})

	r.GET("/contagem-pessoas", func(c *gin.Context) {
		response, status, err := handlers.CountPeopleHandler(db)
		if err != nil {
			c.String(status, "")
			return
		}
		c.String(status, string(response))
	})

	return r
}
