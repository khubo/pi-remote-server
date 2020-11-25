package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khubo/rpi-server/handler"
)

func main() {
	r := gin.Default()
	m := handler.New()
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.Run(":8080")
}
