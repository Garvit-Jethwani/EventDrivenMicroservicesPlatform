package nats

import (
	"fmt"

	"github.com/Garvit-Jethwani/event-gateway-service/models"
	"github.com/nats-io/nats.go"
)

type NATSStore struct {
	conn *nats.Conn
}

func NewNATSStore(url, clusterID string) (*NATSStore, error) {
	var err error
	if clusterID != "" {
		url = fmt.Sprintf("%s://%s", clusterID, url)
	}

	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NATSStore{
		conn: conn,
	}, nil
}

func (n *NATSStore) Publish(event models.Event) (string, error) {
	payload := event.Payload

	msg, err := n.conn.Request("events", payload, 10)
	if err != nil {
		return "", err
	}

	return string(msg.Data), nil
}

func (n *NATSStore) GetEventStatus(eventID string) (models.EventStatus, error) {
	// Implement the logic to retrieve the event status from NATS
	return "", nil
}

func (n *NATSStore) ListEvents(filters ...models.EventFilter) ([]models.EventSummary, error) {
	// Implement the logic to list events from NATS based on filters
	return nil, nil
}

func (n *NATSStore) Close() {
	n.conn.Close()
}
