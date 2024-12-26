package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sultansyah/note-api/internal/config"
	"github.com/sultansyah/note-api/internal/middleware"
	"github.com/sultansyah/note-api/internal/note"
	"github.com/sultansyah/note-api/internal/token"
	"github.com/sultansyah/note-api/internal/user"

	_ "github.com/sultansyah/note-api/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Notes API
// @version         1.0
// @description     A REST API for managing notes and user authentication
// @host           localhost:8080
// @BasePath       /api/v1
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

	fmt.Println("key = ", jwtSecretKey)
	tokenService := token.NewTokenService([]byte(jwtSecretKey))

	userRepository := user.NewUserRepository()
	userService := user.NewUserService(userRepository, db, tokenService)
	userHandler := user.NewUserHandler(userService)

	noteRepository := note.NewNoteRepository()
	noteService := note.NewNoteService(noteRepository, db)
	noteHandler := note.NewNoteHandler(noteService)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")

	api.POST("/auth/register", userHandler.Register)
	api.POST("/auth/login", userHandler.Login)
	api.POST("/auth/name", middleware.AuthMiddleware(tokenService), userHandler.EditName)
	api.POST("/auth/email", middleware.AuthMiddleware(tokenService), userHandler.EditEmail)
	api.POST("/auth/password", middleware.AuthMiddleware(tokenService), userHandler.EditPassword)

	api.POST("/notes", middleware.AuthMiddleware(tokenService), noteHandler.Create)
	api.PUT("/notes/{id}", middleware.AuthMiddleware(tokenService), noteHandler.Edit)
	api.DELETE("/notes/{id}", middleware.AuthMiddleware(tokenService), noteHandler.Delete)
	api.GET("/notes/{id}", middleware.AuthMiddleware(tokenService), noteHandler.FindById)
	api.GET("/notes", middleware.AuthMiddleware(tokenService), noteHandler.FindAll)

	if err := router.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
