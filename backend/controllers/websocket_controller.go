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
func (ctrl *WebSocketController) Connect(c *gin.Context) {
ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
if err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upgrade connection"})
return
}

userID, exists := c.Get("userID")
clientID := "anonymous"
if exists && userID != nil {
if id, ok := userID.(string); ok {
clientID = id
}
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
for msg := range client.Send {
err := client.Conn.WriteJSON(msg)
if err != nil {
return
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
