package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sultansyah/note-api/internal/config"
	"github.com/sultansyah/note-api/internal/note"
	"github.com/sultansyah/note-api/internal/token"
	"github.com/sultansyah/note-api/internal/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	dbConfig := config.DBConfig{
		User:     dbUser,
		Password: dbPassword,
		Host:     dbHost,
		Port:     dbPort,
		Name:     dbName,
	}

	db, err := config.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	tokenService := token.NewTokenService(jwtSecretKey)

	userRepository := user.NewUserRepository()
	userService := user.NewUserService(userRepository, db)
	userHandler := user.NewUserHandler(userService, tokenService)

	noteRepository := note.NewNoteRepository()
	noteService := note.NewNoteService(noteRepository, db)
	noteHandler := note.NewNoteHandler(noteService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.POST("/auth/register", userHandler.Register)
	api.POST("/auth/login", userHandler.Login)
	api.POST("/auth/name", userHandler.EditName)
	api.POST("/auth/email", userHandler.EditEmail)
	api.POST("/auth/password", userHandler.EditPassword)

	api.POST("/notes", noteHandler.Create)
	api.PUT("/notes/{id}", noteHandler.Edit)
	api.DELETE("/notes/{id}", noteHandler.Delete)
	api.GET("/notes/{id}", noteHandler.FindById)
	api.GET("/notes", noteHandler.FindAll)

	if err := router.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
