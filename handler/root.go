package handler

import (
	"encoding/json"
	"fmt"

	"github.com/khubo/rpi-server/mouse"
	"gopkg.in/olahol/melody.v1"
)

//New create and assign a new websocket connection handler
func New() *melody.Melody {

	m := melody.New()

	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("client connected")
	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("client disconnected")
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		var action Action
		json.Unmarshal(msg, &action)

		fmt.Println("action is", action.Device)
		switch action.Device {
		case "mouse":
			mouse.Handler(mouse.Action{action.ActionType, action.X, action.Y})
		default:
			fmt.Println("invalid action")
		}
	})

	return m
}
