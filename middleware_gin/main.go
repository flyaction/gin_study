package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main()  {
	f,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	//r.Use(gin.Logger(),gin.Recovery())
	r.Use(gin.Logger())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name","default name")
		panic("test panic")
		c.String(200,"%",name)
	})
	r.Run()

}
