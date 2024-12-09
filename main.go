package main

import (
	"context"
	"log"
	"os"

	firebaseAdmin "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
	"github.com/kantacky/p2hacks2024-test-api/db"
	api "github.com/kantacky/p2hacks2024-test-api/generated"
	"github.com/kantacky/p2hacks2024-test-api/middleware"
	"github.com/kantacky/p2hacks2024-test-api/server"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if os.Getenv("MODE") != "DEBUG" {
		os.Setenv("GIN_MODE", "release")
	}

	// Migrate
	db.Migrate()

	// Firebase
	app, err := firebaseAdmin.NewApp(context.Background(), nil)
	if err != nil {
		log.Printf("Error initializing app: %v\n", err)
	}

	r := gin.Default()

	// Middleware
	m := middleware.New(app)
	if os.Getenv("MODE") != "DEBUG" {
		r.Use(m.RequireAppCheck())
	}
	r.Use(m.GetFirebaseAuthIDToken())

	// Server
	s := server.New(app)
	api.RegisterHandlers(r, s)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
