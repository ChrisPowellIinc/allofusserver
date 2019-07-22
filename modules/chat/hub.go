// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

import "log"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.email] = client
		case client := <-h.unregister:
			log.Println("Unregister client: ", client.email)
			if _, ok := h.clients[client.email]; ok {
				delete(h.clients, client.email)
				close(client.send)
			}
		case message := <-h.broadcast:
			var (
				okRecipient, okOwner bool
				rec, own             *Client
			)
			recipient, reOk := message["recipient"].(string)
			owner, owOk := message["owner"].(string)
			if !owOk || !reOk {
				log.Println("owner or recipient are not available")
				continue
			}
			rec, okRecipient = h.clients[recipient]
			own, okOwner = h.clients[owner]
			if okRecipient && message["channel"] == "message" {
				log.Printf("Message sent to: %v, From: %v, message: %v\n", recipient, owner, message)
				rec.send <- message
				continue
			} else {
				// log.Println("channel: ", message["channel"])
				// log.Println("recipient: ", message["email"])
				log.Println("there is no recipient to send message to: ", recipient)
			}
			if message["channel"] == "create or join" {
				if !okRecipient {
					// contact the client to accept a call.
					message["channel"] = "error"
					message["error"] = "User is not available"
					own.send <- message
					continue
				} else {
					// contact the client to accept a call.
					message["channel"] = "request"
					rec.send <- message
					continue
				}
			}

			if message["channel"] == "accept" && okRecipient && okOwner {
				// the owner here is the person being contacted.
				// while the recipient is one that contacted this owner.
				message["channel"] = "joined"
				rec.send <- message
				rec.peer = own.email
				own.peer = rec.email
				continue
			}
			if message["channel"] == "decline" && okRecipient {
				message["channel"] = "error"
				message["error"] = "The client declined the call"
				rec.send <- message
				continue
			}
		}
	}
}
