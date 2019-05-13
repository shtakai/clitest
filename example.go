package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type Person2 struct {
	Name string `form:"name"`
	Address string `form:"address"`
}

type User struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
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
	r.GET("/testing", startPage())
	r.POST("/testing2", startPage2())
	r.Run(":9999")
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

func startPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		// If `GET`, only `Form` binding engine (`query`) used.
		// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
		// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
		if c.ShouldBind(&user) == nil {
			log.Println(user.Name)
			log.Println(user.Address)
			log.Println(user.Birthday)
		}
		c.String(http.StatusOK, "success")
	}
}

func startPage2() gin.HandlerFunc {
	return func(c *gin.Context) {
		var person Person2
		name := c.PostForm("name")
		address := c.PostForm("address")
		log.Printf("name: %v   address: %v", name, address)
		if c.ShouldBindQuery(&person) == nil {
			log.Println("=== only bind by query string ===")
			log.Println(person.Name)
			log.Println(person.Address)
		}
		c.String(http.StatusOK, "Success")
	}
}
