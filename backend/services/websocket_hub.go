package services

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketMessage represents a message sent over WebSocket
type WebSocketMessage struct {
	Type       string      `json:"type"`
	Action     string      `json:"action"`
	EntityID   string      `json:"entityId"`
	EntityType string      `json:"entityType"`
	Data       any 			`json:"data"`
	Timestamp  string      `json:"timestamp"`
}

// Client represents a WebSocket client connection
// ID is set to the authenticated user's ID (from JWT)
type Client struct {
	ID   string                        // userID of authenticated user
	Conn *websocket.Conn              // WebSocket connection
	Send chan *WebSocketMessage       // Buffered channel for outgoing messages
}

// Hub maintains active client connections and broadcasts messages
type Hub struct {
	Clients    map[*Client]bool
	BroadcastChan  chan *WebSocketMessage
	Register   chan *Client
	Unregister chan *Client
	mu         sync.RWMutex
}

// NewHub creates a new WebSocket hub
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		BroadcastChan:  make(chan *WebSocketMessage, 256),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run starts the hub and processes messages
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client] = true
			h.mu.Unlock()

		case client := <-h.Unregister:
			h.mu.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
			h.mu.Unlock()

		case message := <-h.BroadcastChan:
			h.mu.RLock()
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					go func(c *Client) {
						h.Unregister <- c
					}(client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// Broadcast sends a message to all connected clients
func (h *Hub) Broadcast(msg *WebSocketMessage) {
	h.BroadcastChan <- msg
}

// GetClientCount returns the number of connected clients
func (h *Hub) GetClientCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.Clients)
}

// NewWebSocketMessage creates a new WebSocket message
func NewWebSocketMessage(msgType, action, entityID, entityType string, data interface{}) *WebSocketMessage {
	return &WebSocketMessage{
		Type:       msgType,
		Action:     action,
		EntityID:   entityID,
		EntityType: entityType,
		Data:       data,
		Timestamp:  getCurrentTimestamp(),
	}
}

// getCurrentTimestamp returns current time in ISO 8601 format
func getCurrentTimestamp() string {
	return time.Now().Format("2006-01-02T15:04:05Z07:00")
}