package main

import (
	"encoding/json"
	"sync"
	"time"

	bolt "go.etcd.io/bbolt"
	"golang.org/x/net/websocket"
)

type server struct {
	ws *websocket.Conn
	sync.Mutex
	db         *bolt.DB
	recipients []string
	serverKey  string
}

type Message struct {
	Action  string          `json:"action,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

type CalibrateData struct {
	Scene int `json:"scene,omitempty"`
}

// LocationData describes location message from user
type LocationData struct {
	UserID string `json:"user_id,omitempty"`
	UUID   string `json:"uuid,omitempty"`
	Minor  int    `json:"minor,omitempty"`
	Major  int    `json:"major,omitempty"`
}

// LocationDataRecord describes location message from user
type LocationDataRecord struct {
	TS    time.Time `json:"ts,omitempty"`
	UUID  string    `json:"uuid,omitempty"`
	Minor int       `json:"minor,omitempty"`
	Major int       `json:"major,omitempty"`
}

// PushNotification describes push message
type PushNotification struct {
	UserID  string `json:"user_id,omitempty"`
	Action  string `json:"action,omitempty"`
	Msg     string `json:"msg,omitempty"`
	PlaceID string `json:"place_id,omitempty"`
	Minor   int    `json:"minor,omitempty"`
	Major   int    `json:"major,omitempty"`
}

func dataToRecord(input LocationData) LocationDataRecord {
	return LocationDataRecord{
		TS:    time.Now(),
		UUID:  input.UUID,
		Minor: input.Minor,
		Major: input.Major,
	}
}

func recordToData(key []byte, input LocationDataRecord) LocationData {
	return LocationData{
		UserID: string(key),
		Minor:  input.Minor,
		Major:  input.Major,
	}
}
