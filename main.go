package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/ws", func(c *gin.Context) {
		fmt.Println("socket connection established")
		m.HandleRequest(c.Writer, c.Request)
	})

	// handle the websocket message
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println("msg is", msg)
	})

	r.Run(":8080")
}
