package main

import "github.com/gin-gonic/gin"
import "github.com/gin-gonic/autotls"

func main()  {

	 r := gin.Default()
	 r.GET("/test", func(c *gin.Context) {
		 c.String(200,"hello other auto tls")
	 })
	 autotls.Run(r,"www.itpp.tk")
}


