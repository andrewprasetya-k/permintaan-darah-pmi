package controllers

import (
	"backend/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	hub *services.Hub
}

func NewWebSocketController(hub *services.Hub) *WebSocketController {
	return &WebSocketController{hub: hub}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Connect handles WebSocket connection
// Token must be provided as query parameter (?token=JWT_TOKEN) because
// HTTP protocol doesn't allow custom headers during WebSocket upgrade.
// The JWTMiddleware (applied to this route) validates the token and extracts userID
func (ctrl *WebSocketController) Connect(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upgrade connection"})
		return
	}

	// userID is set by JWTMiddleware after validating the token from query param
	userID, exists := c.Get("userID")
	if !exists || userID == nil {
		// This shouldn't happen because JWTMiddleware would have rejected the request
		ws.Close()
		return
	}

	clientID, ok := userID.(string)
	if !ok {
		ws.Close()
		return
	}

	client := &services.Client{
		ID:   clientID,
		Conn: ws,
		Send: make(chan *services.WebSocketMessage, 256),
	}

	ctrl.hub.Register <- client

	go ctrl.handleRead(client, ctrl.hub)
	go ctrl.handleWrite(client)
}

// handleRead reads messages from WebSocket client
func (ctrl *WebSocketController) handleRead(client *services.Client, hub *services.Hub) {
	defer func() {
		hub.Unregister <- client
		client.Conn.Close()
	}()

	// Set longer timeout: 24 hours (keep alive during admin session)
	client.Conn.SetReadDeadline(time.Now().Add(24 * time.Hour))
	// Set pong handler to reset read deadline on pong messages
	client.Conn.SetPongHandler(func(string) error {
		client.Conn.SetReadDeadline(time.Now().Add(24 * time.Hour))
		return nil
	})

	for {
		var msg services.WebSocketMessage
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// Log error if needed
			}
			break
		}

		msg.Timestamp = time.Now().Format("2006-01-02T15:04:05Z07:00")
		hub.Broadcast(&msg)
	}
}

// handleWrite writes messages to WebSocket client
func (ctrl *WebSocketController) handleWrite(client *services.Client) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case msg := <-client.Send:
			client.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err := client.Conn.WriteJSON(msg)
			if err != nil {
				return
			}
		case <-ticker.C:
			// Send ping to keep connection alive
			client.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := client.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// GetConnectionStatus returns current WebSocket connections info
func (ctrl *WebSocketController) GetConnectionStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success":           true,
		"activeConnections": ctrl.hub.GetClientCount(),
	})
}