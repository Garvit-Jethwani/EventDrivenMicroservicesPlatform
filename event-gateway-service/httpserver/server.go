package httpserver

import (
	"net/http"

	"github.com/Garvit-Jethwani/event-gateway-service/events"
)

type Server struct {
	addr    string
	handler *events.EventHandler
}

func NewServer(addr string, handler *events.EventHandler) *Server {
	return &Server{
		addr:    addr,
		handler: handler,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/events", s.handler.IngestEvent)
	http.HandleFunc("/events/", s.handleGetEvent)
	http.HandleFunc("/events", s.handleListEvents)

	return http.ListenAndServe(s.addr, nil)
}

func (s *Server) handleGetEvent(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle GET /events/{eventId}
}

func (s *Server) handleListEvents(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle GET /events
}
