package events

import (
	"encoding/json"
	"net/http"

	"github.com/Garvit-Jethwani/event-gateway-service/models"
)

type EventHandler struct {
	eventStore EventStore
}

func NewEventHandler(eventStore EventStore) *EventHandler {
	return &EventHandler{
		eventStore: eventStore,
	}
}

func (h *EventHandler) IngestEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventID, err := h.eventStore.Publish(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(eventID))
}

// Implement other handlers for GET /events/{eventId} and GET /events
