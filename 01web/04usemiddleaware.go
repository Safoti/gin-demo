package main

import "github.com/gin-gonic/gin"

/**

  使用中间件
 */
/**
  这一块将在后面具体实现 请先忽略
*/
func main() {
	// 创建一个不包含中间件的路由器
	r := gin.New()

	// 全局中间件
	// 使用 Logger 中间件
	r.Use(gin.Logger())
	// 使用 Recovery 中间件
	r.Use(gin.Recovery())
	// 路由添加中间件，可以添加任意多个
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	// 路由组中添加中间件
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:

	//authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)
	//
	//	// nested group
	//	testing := authorized.Group("testing")
	//	testing.GET("/analytics", analyticsEndpoint)
	//}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
