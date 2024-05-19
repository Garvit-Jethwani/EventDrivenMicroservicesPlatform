package main

import (
	"log"
	"strings"

	"github.com/Garvit-Jethwani/notification-service/config"
	"github.com/Garvit-Jethwani/notification-service/events"
	"github.com/Garvit-Jethwani/notification-service/notification"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	brokers := strings.Split(cfg.EventStoreConfig["brokers"], ",")
	consumer, err := events.NewKafkaConsumer(brokers, cfg.EventStoreConfig["group.id"], cfg.EventStoreConfig["topic"])
	defer consumer.Close()

	// Create notification service
	notificationService := notification.NewNotificationService(consumer)

	// Start the notification service
	err = notificationService.Start()
	if err != nil {
		log.Fatalf("Failed to start notification service: %v", err)
	}
}
