package main

import "github.com/gin-gonic/gin"


var r *gin.Engine

func main() {
	r = gin.Default()


	initRouter()

	r.Run() // listen and server on 0.0.0.0:8080
}

func initRouter()  {
	r.GET("/ping",pingHandler)

	r.GET("/ping2",pingHandler2)
	r.GET("/ping3",pingHandler3)

	r.GET("/service/item",serviceItem)
	r.GET("/service/list",serviceList)

	r.GET("/user/get",GetUserInfo)
}

func  pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func  pingHandler2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong2",
	})
}

func  pingHandler3(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong3",
	})
}


func serviceItem(c *gin.Context)  {


	c.JSON(200, gin.H{
		"message": "servicItem",
	})

}
func serviceList(c *gin.Context)  {



	c.JSON(200, gin.H{
		"message": "serviceList",
	})

}

func GetUserInfo(c *gin.Context)  {
	c.JSON(200, gin.H{
		"name": "可润",
		"sex":"女",
	})
}