package app_socket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocket struct {
	Conn   *websocket.Conn
	Out    chan []byte
	In     chan []byte
	Events map[string]EventHandler
}

func NewWebSocket(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("An error occured while upgrading the connection %v", err)
	}
	ws := &WebSocket{
		Conn:   conn,
		Out:    make(chan []byte),
		In:     make(chan []byte),
		Events: make(map[string]EventHandler),
	}

	// go ws.Reader()
	go ws.Writer()
	return ws, nil
}

func (ws *WebSocket) Reader() {
	defer func() {
		ws.Conn.Close()
	}()

	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Websocket message error: %v", err)
			}
			break
		}
		event, err := NewEventFromRaw(message)
		if err != nil {
			log.Printf("Error parsing message: %v", err)
		} else {
			log.Printf("Message : %v", event)
		}

		if action, ok := ws.Events[event.Name]; ok {
			action(event)
		}
	}
}

func (ws *WebSocket) Writer() {
	for i := 0; i < 1000; i++ {
		ws.Conn.WriteMessage(websocket.CloseMessage, make([]byte, 1))
		return
	}
	w, err := ws.Conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	w.Write([]byte("HELLOW JING"))
	w.Close()
}

func (ws *WebSocket) On(eventName string, action EventHandler) *WebSocket {
	ws.Events[eventName] = action
	return ws
}
