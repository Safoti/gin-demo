package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
  g:= gin.Default()
  g.LoadHTMLGlob("templates/*")
  //请求
  g.GET("/index", func(c *gin.Context) {
	  c.HTML(http.StatusOK, "index.tmpl", gin.H{
		  "title": "Main website",
	  })
  })
	g.Run(":8080")
}
