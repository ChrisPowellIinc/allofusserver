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
			rec, ok_recipient := h.clients[message.Email]
			own, ok_owner := h.clients[message.Owner]
			if ok_recipient && message.Channel == "message" {
				log.Printf("Message sent to: %v, From: %v, message: %v\n", message.Email, message.Owner, message)
				rec.send <- message
				continue
			} else {
				log.Println("channel: ", message.Channel)
				log.Println("recipient: ", message.Email)
				log.Println("type: ", message.Type)
				// log.Println("message: ", message)
			}
			if message.Channel == "create or join" {
				if !ok_recipient {
					// contact the client to accept a call.
					message.Channel = "error"
					message.Error = "User is not available"
					own.send <- message
					continue
				} else {
					// contact the client to accept a call.
					message.Channel = "request"
					rec.send <- message
					continue
				}
			}

			if message.Channel == "accept" && ok_recipient && ok_owner {
				// the owner here is the person being contacted.
				// while the recipient is one that contacted this owner.
				message.Channel = "joined"
				rec.send <- message
				rec.peer = own.email
				own.peer = rec.email
				continue
			}
			if message.Channel == "decline" && ok_recipient {
				message.Channel = "error"
				message.Error = "The client declined the call"
				rec.send <- message
				continue
			}
		}
	}
}
