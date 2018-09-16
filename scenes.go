package main

import (
	"fmt"
)

func (s *server) checkScene(msg CalibrateData) ([]PushNotification, error) {
	records, err := s.readAllDB()
	if err != nil {
		return nil, err
	}

	l := len(records)
	if l < 10 {
		return nil, fmt.Errorf("DB includes only %d records", l)
	}

	result := []PushNotification{}

	splitRec := split(records)

	numFirst := l / 5

	// move to first class
	if len(splitRec[31]) < numFirst {
		diff := numFirst - len(splitRec[31])
		cur32 := splitRec[32]
		cur31 := splitRec[31]
		for i := 0; i < diff; i++ {
			cur31 = append(cur31, cur32[i])
			result = append(result, PushNotification{
				UserID: cur32[i].UserID,
				Action: "move_first",
				Msg:    "Please proceed to first class, your ticket was upgraded",
				Minor:  31,
				Major:  cur32[i].Major,
			})
		}
		splitRec[31] = cur31
		cur32 = append(cur32[:diff-1], cur32[diff:]...)
		splitRec[32] = cur32
	}

	// move to different coach of second class
	numSecond := (len(splitRec[32]) + len(splitRec[33])) / 2
	if len(splitRec[33]) < numSecond {
		diff := numSecond - len(splitRec[33])
		cur32 := splitRec[32]
		cur33 := splitRec[33]
		for i := 0; i < diff; i++ {
			cur33 = append(cur33, cur32[i])
			result = append(result, PushNotification{
				UserID: cur32[i].UserID,
				Action: "move_second",
				Msg:    "Please proceed to different coach",
				Minor:  33,
				Major:  cur32[i].Major,
			})
		}
		splitRec[33] = cur33
		cur32 = append(cur32[:diff-1], cur32[diff:]...)
		splitRec[32] = cur32
	} else {
		diff := numSecond - len(splitRec[32])
		cur32 := splitRec[32]
		cur33 := splitRec[33]
		for i := 0; i < diff; i++ {
			cur32 = append(cur32, cur33[i])
			result = append(result, PushNotification{
				UserID: cur33[i].UserID,
				Action: "move_second",
				Msg:    "Please proceed to different coach",
				Minor:  32,
				Major:  cur33[i].Major,
			})
		}
		splitRec[32] = cur32
		cur33 = append(cur33[:diff-1], cur33[diff:]...)
		splitRec[33] = cur33
	}

	return result, nil
}

func split(records []LocationData) map[int][]LocationData {
	result := map[int][]LocationData{
		31: []LocationData{},
		32: []LocationData{},
		33: []LocationData{},
	}
	for _, rec := range records {
		cur := result[rec.Minor]
		cur = append(cur, rec)
		result[rec.Minor] = cur
	}
	return result
}
