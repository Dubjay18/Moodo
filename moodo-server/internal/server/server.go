package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"firebase.google.com/go/auth"
	_ "github.com/joho/godotenv/autoload"

	"moodo-server/internal/database"
	"moodo-server/internal/lib"
)

type Server struct {
	port       int
	authClient *auth.Client
	db         database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	firebaseApp := lib.InitializeFirebaseApp()
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	NewServer := &Server{
		port:       port,
		authClient: authClient,
		db:         database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
