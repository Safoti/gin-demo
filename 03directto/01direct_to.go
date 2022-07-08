package main

import "github.com/gin-gonic/gin"

func main() {
	r:=gin.Default()
	//test --->test2
	r.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	r.Run(":8080")
}
