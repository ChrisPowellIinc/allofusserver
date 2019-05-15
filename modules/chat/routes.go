package chat

import (
	"log"
	"net/http"

	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	"github.com/ChrisPowellIinc/allofusserver/internal/jwt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

// Handler : Routes handler
type Handler struct {
	config *config.Config
}

var handler *Handler

// New : Creates a new handler object
func New(config *config.Config) *Handler {
	return &Handler{config: config}
}

// MyCorsServer Hold the server instance to be called after the cors is set.
// type MyCorsServer struct {
// 	server *socketio.Server
// }

// func (s *MyCorsServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
// 	// log.Println("You are an idiot")
// 	// log.Println(req.Header.Get("Origin"))
// 	// log.Println(req.Method)

// 	if origin := req.Header.Get("Origin"); origin != "" {
// 		rw.Header().Set("Access-Control-Allow-Origin", origin)
// 		rw.Header().Set("Access-Control-Allow-Credentials", "true")
// 		rw.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE, OPTIONS")
// 		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
// 	}
// 	if req.Method == "OPTIONS" {
// 		log.Println("OPTONS GO")
// 		return
// 	}
// 	s.server.ServeHTTP(rw, req)
// }

type myWebsocketServer struct {
	hub *Hub
}

var myWs = myWebsocketServer{
	hub: newHub(),
}

func (ws myWebsocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving websocket from here...")
	serveWs(ws.hub, w, r)
}

// Routes : Defines API routes for this module
func Routes(config *config.Config) http.Handler {
	handler = New(config)
	router := chi.NewRouter()
	router.Use(jwtauth.Verifier(jwt.TokenAuth))
	router.Use(jwt.AuthHandler)

	go myWs.hub.run()
	router.Mount("/", myWs)
	return router
}
