package hub

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/catielanier/hitbox-input-viewer/internal/controller"
	"github.com/gorilla/websocket"
)

type Hub struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]chan []byte
	latest  []byte
}

func New() *Hub {
	initial, _ := json.Marshal(map[string]any{
		"type":  "state",
		"state": controller.Empty(),
	})

	return &Hub{
		clients: make(map[*websocket.Conn]chan []byte),
		latest:  initial,
	}
}

func (h *Hub) Publish(state controller.State) {
	message, _ := json.Marshal(map[string]any{
		"type":  "state",
		"state": state,
	})

	h.mu.Lock()
	defer h.mu.Unlock()

	h.latest = message

	for _, outbound := range h.clients {
		// Only the latest snapshot matters.
		// Never block the input reader on a slow browser.
		select {
		case outbound <- message:
		default:
			select {
			case <-outbound:
			default:
			}

			select {
			case outbound <- message:
			default:
			}
		}
	}
}

func (h *Hub) ServeWS(w http.ResponseWriter, r *http.Request) {
	upgrade := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")

			return origin == "" ||
				origin == "http://"+r.Host ||
				origin == "https://"+r.Host
		},
	}

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	outbound := make(chan []byte, 1)

	h.mu.Lock()
	h.clients[conn] = outbound
	outbound <- h.latest
	h.mu.Unlock()

	defer func() {
		h.mu.Lock()
		delete(h.clients, conn)
		h.mu.Unlock()

		conn.Close()
	}()

	// Keep reading control frames so close/ping/pong work.
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()

	ping := time.NewTicker(25 * time.Second)
	defer ping.Stop()

	for {
		select {
		case msg := <-outbound:
			conn.SetWriteDeadline(
				time.Now().Add(5 * time.Second),
			)

			if err := conn.WriteMessage(
				websocket.TextMessage,
				msg,
			); err != nil {
				return
			}

		case <-ping.C:
			conn.SetWriteDeadline(
				time.Now().Add(5 * time.Second),
			)

			if err := conn.WriteMessage(
				websocket.PingMessage,
				nil,
			); err != nil {
				return
			}

		case <-done:
			return
		}
	}
}
