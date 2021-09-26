package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(20*time.Second)
		c.String(200,"hello other shutdown")
	})


	srv := &http.Server{
		Addr: ":8085",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe();err != nil && err != http.ErrServerClosed{
			log.Fatalf("listen：%s\n",err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	log.Println("shutdown server ...")

	ctx,cancel := context.WithTimeout(context.Background(),20*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx);err != nil{
		log.Fatalf("server shutdown:",err)
	}

	log.Println("server exiting")


}
