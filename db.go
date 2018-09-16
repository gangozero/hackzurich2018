package main

import (
	"encoding/json"
	"log"

	bolt "go.etcd.io/bbolt"
)

func newDB() *bolt.DB {
	db, err := bolt.Open("hack.db", 0600, nil)
	if err != nil {
		log.Fatalf("Can't open DB: %s", err.Error())
	}

	// Setup the location bucket.
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("location"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Can't create bucket: %s", err.Error())
	}

	return db
}

func (s *server) saveLogDB(msg LocationData) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		key := []byte(msg.UserID)
		valueStr := dataToRecord(msg)
		value, err := json.Marshal(valueStr)
		if err != nil {
			return err
		}

		b := tx.Bucket([]byte("location"))
		return b.Put(key, value)
	})
	return err
}

func (s *server) readAllDB() ([]LocationData, error) {
	result := []LocationData{}
	err := s.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("location"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var val LocationDataRecord
			err := json.Unmarshal(v, &val)
			if err != nil {
				return nil
			}
			result = append(result, recordToData(k, val))
		}

		err := tx.DeleteBucket([]byte("location"))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte("location"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
