package main

import "github.com/gin-gonic/gin"

/**
    路由分组
 */
func main() {
	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func readEndpoint(context *gin.Context) {
	context.JSON(200,gin.H{
		"code":222,
	})
}

func submitEndpoint(context *gin.Context) {
	context.JSON(200,gin.H{
		"code":222,
	})
}

func loginEndpoint(context *gin.Context) {
	context.JSON(200,gin.H{
		"code":222,
	})
}
