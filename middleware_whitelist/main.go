package main

import "github.com/gin-gonic/gin"


func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag  := false

		clientIp := c.ClientIP()
		for _,host := range ipList{
			if clientIp == host{
				flag = true
				break;
			}
		}

		if !flag{
			c.String(401,"%s,not in ipList",clientIp)
			c.Abort()
		}
	}
}

func main()  {
	r := gin.Default()

	r.Use(IPAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		c.String(200,"hello middleware_whitelist")
	})

	r.Run()

}