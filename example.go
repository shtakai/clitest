package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/users/:name/:id", userInfo())
	r.GET("/ping", ping())
	r.GET("/somejson", somejson())
	r.Run()
}

func userInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	}
}

func somejson() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO言語",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	}
}

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
