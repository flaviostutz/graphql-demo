package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func makeRoutes(todo Todo) http.Handler {
	ok := func(c *gin.Context) {
		c.String(200, "")
	}

	cors := func(c *gin.Context) {
		c.Writer.Header().Add("access-control-allow-origin", "*")
		c.Writer.Header().Add("access-control-allow-headers", "accept, content-type")
		c.Writer.Header().Add("access-control-allow-methods", "GET,HEAD,POST,DELETE,OPTIONS,PUT,PATCH")
	}

	routes := gin.Default()
	routes.Use(cors)

	routes.OPTIONS("/todos", ok)

	routes.OPTIONS("/todos/:id", ok)

	routes.GET("/todos", func(c *gin.Context) {
		c.JSON(200, todo.All())
	})

	routes.GET("/todos/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		a := todo.Find(id)
		if a == nil {
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, a)
	})

	routes.POST("/todos", func(c *gin.Context) {
		nitem := TodoItem{}
		err := c.BindJSON(&nitem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Couldn't parse body contents. err=%s", err)})
			return
		}
		if nitem.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("'title' is required")})
			return
		}
		item := todo.Create(nitem)
		c.Writer.Header().Add("Location", os.Getenv("BASE_LOCATION_URL")+"/todos/"+item.ID)
		c.JSON(201, item)
	})

	routes.PATCH("/todos/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		item := todo.Find(id)
		if item == nil {
			c.JSON(404, gin.H{})
			return
		}
		err := c.BindJSON(&item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Couldn't parse body contents. err=%s", err)})
			return
		}
		c.JSON(200, gin.H{"message": "Updated successfully"})
	})

	routes.DELETE("/todos", func(c *gin.Context) {
		todo.DeleteAll()
		c.JSON(200, gin.H{"message": "Deleted successfully"})
	})

	routes.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		item := todo.Find(id)
		if item == nil {
			c.JSON(404, gin.H{})
			return
		}
		todo.Delete(id)
		c.JSON(200, gin.H{"message": "Deleted successfully"})
	})

	return routes
}
