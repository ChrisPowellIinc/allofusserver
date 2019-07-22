// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ChrisPowellIinc/allofusserver/models"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 2048
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan Message

	// the clients email address to identify her by...
	email string

	// peer to send message to.
	peer string
}

// Message defines the way messages are exchanged
// type Message struct {
// 	Type    string `json:"type"`
// 	Email   string `json:"email"`
// 	Message string `json:"message"`
// 	Channel string `json:"channel"`
// 	Error   string `json:"error"`
// 	Owner   string `json:"owner"`
// 	// the call candidate details...
// 	Label     int    `json:"label"`
// 	ID        string `json:"id"`
// 	Candidate string `json:"candidate"`
// }
type Message map[string]interface{}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	// c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			log.Println(err)
			break
		}
		m := Message{}
		err = json.Unmarshal(message, &m)
		if err != nil {
			log.Println("Invalid message", err)
			// w.Write([]byte(`{"error": "Invalid Message"}`))
			continue
		}
		// message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- m
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			mm, err := json.Marshal(message)
			if err != nil {
				log.Println("Invalid message", err)
				w.Write([]byte(`{"error": "Invalid Message"}`))
				continue
			}
			w.Write(mm)

			log.Printf("Write(%v) (message) \n", message)
			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				log.Printf("Write(\\n:%#v) newline \n", newline)
				m := <-c.send

				mm, err := json.Marshal(m)
				if err != nil {
					log.Println("Invalid message: ", err)
					w.Write([]byte(`{"error": "Invalid Message"}`))
					continue
				}

				log.Printf("Write(m:%#v) newline \n", string(mm))
				w.Write(mm)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("Error pinging the connection.")
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true /* check the origin here...*/ }
	// r.URL.Query().Get("jwt")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade Error: ", err)
		return
	}

	email, err := models.GetLoggedInUserID(r.Context())
	if err != nil {
		log.Println("Getting logged in user email: ", err)
		return
	}
	log.Println("New Client: ", email)
	client := &Client{
		hub:   hub,
		conn:  conn,
		send:  make(chan Message),
		email: email,
	}
	client.hub.register <- client

	log.Println("Client Added")
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
