package main

import (
	"fmt"
	"log"

	"github.com/edganiukov/fcm"
)

func pushToFCM(msg PushNotification) map[string]interface{} {
	return map[string]interface{}{
		"action":   msg.Action,
		"msg":      msg.Msg,
		"place_id": msg.PlaceID,
		"minor":    msg.Minor,
		"major":    msg.Major,
	}
}

func (s *server) sendPush(msgs []PushNotification) error {
	if len(msgs) == 0 {
		return fmt.Errorf("Push notification can't be empty")
	}

	msg := msgs[0]

	fcmMsg := &fcm.Message{
		RegistrationIDs: s.recipients,
		Data:            pushToFCM(msg),
		Notification: &fcm.Notification{
			Title: msg.Msg,
			Body:  "",
		},
	}

	client, err := fcm.NewClient(s.serverKey)
	if err != nil {
		return err
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(fcmMsg)
	log.Printf("[DEBUG] FCM response: %+v", response)
	return err
}
