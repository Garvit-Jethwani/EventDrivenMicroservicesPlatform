package models

import "time"

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Payload   []byte    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}
