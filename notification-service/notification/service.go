package notification

import (
	"github.com/Garvit-Jethwani/notification-service/events"
	"github.com/Garvit-Jethwani/notification-service/models"
)

type NotificationService struct {
	consumer events.EventConsumer
}

func NewNotificationService(consumer events.EventConsumer) *NotificationService {
	return &NotificationService{
		consumer: consumer,
	}
}

func (n *NotificationService) Start() error {
	eventChan, err := n.consumer.Consume()
	if err != nil {
		return err
	}

	for event := range eventChan {
		err = n.processEvent(event)
		if err != nil {
			// Handle error
		}
	}

	return nil
}

func (n *NotificationService) processEvent(event models.Event) error {
	// Implement the logic to send notifications based on the event
	return nil
}
