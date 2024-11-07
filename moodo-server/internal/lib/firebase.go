package lib

import (
	"context"
	"firebase.google.com/go"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/api/option"
	"log"
	"os"
)

var firebaseServiceKey = os.Getenv("FIREBASE_SERVICE_ACCOUNT_KEY_PATH")

func InitializeFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile(firebaseServiceKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}
	return app
}
