package main

import (
	"encoding/json"
	"log"

	"golang.org/x/net/websocket"
)

func (s *server) handlerWS(ws *websocket.Conn) {
	log.Printf("[INFO] New connection, update server")
	s.Lock()
	s.ws = ws
	s.Unlock()
	log.Printf("[INFO] Serve")

	for {
		if !s.ws.IsServerConn() {
			log.Printf("[INFO] Dead connection, close")
			s.ws.Close()
			return
		}
		var msg Message
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			log.Printf("[ERROR] Error read location from WS: %s", err.Error())
			log.Printf("[ERROR] Dead connection, close")
			s.ws.Close()
			return
		}

		switch msg.Action {
		case "location":
			s.doLocation(msg.Payload)
		case "calibrate":
			s.doCalibrate(msg.Payload)
		}
	}
}

func (s *server) doLocation(raw json.RawMessage) {
	var msg LocationData
	err := json.Unmarshal(raw, &msg)
	if err != nil {
		log.Printf("[ERROR] Can't decode LocationData: %s", err.Error())
	}
	// log to CLI
	checkUserAndLog(msg)

	// save to DB
	err = s.saveLogDB(msg)
	if err != nil {
		log.Printf("[ERROR] Can't save location to DB: %s", err.Error())
	}
}

func checkUserAndLog(msg LocationData) {
	if msg.UserID == "" {
		log.Printf("[ERROR] no user_id in message '%+v'", msg)
		return
	}
	log.Printf("New location: user '%s' -> %d::%d\n", msg.UserID, msg.Major, msg.Minor)
}

func (s *server) doCalibrate(raw json.RawMessage) {
	var msg CalibrateData
	err := json.Unmarshal(raw, &msg)
	if err != nil {
		log.Printf("[ERROR] Can't decode CalibrateData: %s", err.Error())
	}
	// check demo
	push, err := s.checkScene(msg)
	if err != nil {
		log.Printf("[ERROR] Can't check scene: %s", err.Error())
	}

	if push != nil {
		err = websocket.JSON.Send(s.ws, push)
		if err != nil {
			log.Printf("[ERROR] Error write push to WS: %s", err.Error())
		}

		err = s.sendPush(push)
		if err != nil {
			log.Printf("[ERROR] Error send push through FCM: %s", err.Error())
		}
	}
}
