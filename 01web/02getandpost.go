package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
  g:=gin.Default()
  //普通的get.post请求
  //get 请求
  g.GET("/get", func(ctx *gin.Context) {
     ctx.JSON(200,gin.H{
     	"message":"get",
     	"code":200,
	 })

  })
  //post 请求
  g.POST("/post", func(ctx *gin.Context) {
	  ctx.JSON(200, gin.H{
		  "message": "post",
	  })
  })



  //get方式：请求url参数
//http://localhost:8080/get/3?name=lktbz&pwd=zk 对应的请求路径
//当不传其中的一个参数http://localhost:8080/get/3?name=lktbz
/**
  {
      "code": 200,
      "id": "3",
      "message": "获取请求路径中的参数",
      "name": "lktbz",
      "pwd": ""
  }
  返回结果如下  ：解决办法，给个默认值
 */
	g.GET("/get/:id", func(ctx *gin.Context) {
		id:=ctx.Param("id")  //url
		name:=ctx.Query("name") //请求参数
		//pwd:=ctx.Query("pwd") //请求参数
		pwd:=ctx.DefaultQuery("pwd","default")
		ctx.JSON(200,gin.H{
			"id":id,
			"name":name,
			"code":200,
			"message":"获取请求路径中的参数",
			"pwd":pwd,
		})

	})
	// 但是，这个规则既能匹配/get/john/格式也能匹配/get/john/send这种格式
	// 如果没有其他路由器匹配/get/john，它将重定向到/get/john/
	g.GET("/get/:id/*action", func(ctx *gin.Context) {
		id:=ctx.Param("id")  //url
		name:=ctx.Query("name") //请求参数
		//pwd:=ctx.Query("pwd") //请求参数
		pwd:=ctx.DefaultQuery("pwd","default")
		ctx.JSON(200,gin.H{
			"id":id,
			"name":name,
			"code":200,
			"message":"获取请求路径中的参数",
			"pwd":pwd,
		})

	})


	//post方式请求传参
	g.POST("/posts", func(ctx *gin.Context) {
		name:=ctx.DefaultPostForm("name","dn")
		pwd:=ctx.DefaultPostForm("pwd","dp")
		ctx.JSON(200,gin.H{
			"name":name,
			"code":200,
			"message":"获取请求路径中的参数",
			"pwd":pwd,
		})

	})



	//Get + Post 混合
//POST /post?id=1234&page=1 HTTP/1.1
	//Content-Type: application/x-www-form-urlencoded
	//
	//name=manu&message=this_is_great
	g.POST("/postg", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
  g.Run(":8080")
}
