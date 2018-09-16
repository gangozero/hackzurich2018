package main

import (
	"log"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

func newServer() *server {
	return &server{
		serverKey:  os.Getenv("SERVER_KEY"),
		recipients: []string{os.Getenv("TEST_DEVICE_TOKEN")},
	}
}

func (s *server) ping() {
	for {
		if s.ws != nil && s.ws.IsServerConn() {
			err := websocket.Message.Send(s.ws, "ping")
			if err != nil {
				log.Printf("[ERROR] Can't PING the client %s", err.Error())
			}
		} else {
			log.Println("[INFO] No WS clients connected")
		}

		time.Sleep(10 * time.Second)
	}

}
