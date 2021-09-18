package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.Static("/assets","./assets")  // go build -o router_static && ./router_static
	r.StaticFS("static",http.Dir("static"))
	r.StaticFile("/favicon.ico","./favicon.ico")
	r.Run()

}
