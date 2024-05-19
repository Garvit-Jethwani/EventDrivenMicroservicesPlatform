package models

import "time"

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Payload   []byte    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

type EventStatus string

const (
	EventStatusPending   EventStatus = "PENDING"
	EventStatusProcessed EventStatus = "PROCESSED"
	EventStatusFailed    EventStatus = "FAILED"
)

type EventSummary struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Status    EventStatus `json:"status"`
	Timestamp time.Time   `json:"timestamp"`
}

type EventFilter struct {
	Type      string
	StartTime time.Time
	EndTime   time.Time
}
