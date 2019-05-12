package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC string `form:"field_c"`
}

type StructD struct {
	NestedAnonymousStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}


func main() {
	r := gin.Default()
	r.GET("/users/:name/:id", userInfo())
	r.GET("/ping", ping())
	r.GET("/somejson", somejson())
	r.GET("/getB", GetDataB())
	r.GET("/getC", GetDataC())
	r.GET("/getD", GetDataD())
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

func GetDataB() gin.HandlerFunc {
	return func(c *gin.Context) {
		var b StructB
		c.Bind(&b)
		c.JSON(200, gin.H{
			"a": b.NestedStruct,
			"b": b.FieldB,
		})
	}
}

func GetDataC() gin.HandlerFunc {
	return func(c *gin.Context) {
		var structC StructC
		c.Bind(&structC)
		c.JSON(200, gin.H{
			"a": structC.NestedStructPointer,
			"c": structC.FieldC,
		})
	}
}

func GetDataD() gin.HandlerFunc {
	return func(c *gin.Context) {
		var d StructD
		c.Bind(&d)
		c.JSON(200, gin.H{
			"x": d.NestedAnonymousStruct,
			"d": d.FieldD,
		})
	}
}
