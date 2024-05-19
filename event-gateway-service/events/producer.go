package events

import "github.com/Garvit-Jethwani/event-gateway-service/models"

type EventStore interface {
	Publish(event models.Event) (string, error)
	GetEventStatus(eventID string) (models.EventStatus, error)
	ListEvents(filters ...models.EventFilter) ([]models.EventSummary, error)
	Close()
}
