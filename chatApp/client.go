package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user.

type client struct {

	socket *websocket.Conn
	receive chan *message
	room *room
    userData map[string]interface{}
}
func (c *client) read() {
	for {
	var msg *message
	if err := c.socket.ReadJSON(&msg); err == nil {
	msg.When = time.Now()
	msg.Name = c.userData["name"].(string)
	c.room.forward <- msg
	} else {
	break
	}
	}
	c.socket.Close()
	}
func (c *client) write() {
	for msg := range c.receive {
	if err := c.socket.WriteJSON(msg); err != nil {
	break
	}
	}
	c.socket.Close()
	}