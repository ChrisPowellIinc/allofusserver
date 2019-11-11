package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"

	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	"github.com/ChrisPowellIinc/allofusserver/internal/jwt"
	logger "github.com/ChrisPowellIinc/allofusserver/internal/log"
	"github.com/ChrisPowellIinc/allofusserver/migrations"
	"github.com/ChrisPowellIinc/allofusserver/modules/auth"
	"github.com/ChrisPowellIinc/allofusserver/modules/chat"
	"github.com/ChrisPowellIinc/allofusserver/modules/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/rs/cors"
)

// Routes : Registers all available routes from modules
func Routes(config *config.Config) *chi.Mux {
	v := chi.NewRouter()
	v.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		// middleware.RedirectSlashes,         // Redirect slashes to no slash URL versions
		middleware.Recoverer,               // Recover from panics without crashing server
		middleware.Timeout(60*time.Second), // Timeout requests after 60 seconds
	)

	chiCors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "CONNECT"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Content-Type", "*"},
		Debug:            false,
	})
	v.Use(chiCors.Handler)

	// set up handler for Websocket.
	v.Mount("/socket.io/", chat.Routes(config))

	v.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(apiRouter chi.Router) {
			apiRouter.Mount("/auth", auth.Routes(config))
			apiRouter.Mount("/users", user.Routes(config))
		})

	})

	filesDir := filepath.Join("public/assets")
	fileServer := http.StripPrefix("/assets/",
		http.FileServer(http.Dir(filesDir)))
	v.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		fileServer.ServeHTTP(w, r)
	})

	v.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello there, i am still testing the default route")
		http.ServeFile(w, r, "./public/index.html")
	})

	v.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})
	return v
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var isTest = flag.Bool("test", false, "Set test mode to use mock io resources")
	var isDebug = flag.Bool("debug", false, "Set Debug mode to print config data")
	var isMigrations = flag.Bool("migrations", false, "Run migrations against the database")
	flag.Parse()

	config := config.GetConf(*isTest, *isDebug)
	defer config.DBSession.Close()

	if *isMigrations {
		migrations.MakeMigrations(config)
		log.Println("Successfully ran migrations")
		return
	}

	jwt.Register([]byte(config.JWTSecret))
	logger.Register(config.Constants.LogDir)

	router := Routes(config)

	walkFunc := func(method string, route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("👉  %s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("⚠️  Logging err: %s\n", err.Error())
	}

	// This block of code will allow graceful shutdown of our server,
	// using the server Shurdown method which is a part lf the standard library
	PORT := ":"
	if !*isDebug && !*isTest {
		PORT += os.Getenv("API_PORT")
	} else {
		PORT += config.Constants.PORT
	}
	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Println("😔 Shutting down. Goodbye...")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("⚠️  HTTP server Shutdown error: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Printf("Serving at 🔥 %s \n", server.Addr)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("⚠️  HTTP server ListenAndServe error: %v", err)
	}

	<-idleConnsClosed
}
