# Notes API

REST API for managing notes with user authentication built using Go and Gin framework.

This is the backend for the Note API project. For the frontend, you can access it through the link sultansyah/note-frontend.

## Features
- User authentication with JWT
- CRUD operations for notes
- MySQL database
- Swagger API documentation

## Prerequisites
- Go
- MySQL
- migrate CLI tool

## Setup Instructions

### 1. Clone Repository
```bash
git clone https://github.com/sultansyah/note-api.git
cd note-api
```

### 2. Environment Variables
Copy `.env.example` to `.env`:
```bash
cp .env.example .env
```

Update the values in `.env`:
```env
JWT_SECRET_KEY=your_jwt_secret
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=localhost
DB_PORT=3306
```

### 3. Database Migration
Run MySQL migrations:
```bash
migrate -database "mysql://username:password@tcp(localhost:3306)/database_name" -path database/migrations up
```

### 4. Install Dependencies
```bash
go mod tidy
```

### 5. Generate Swagger Documentation
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

### 6. Run Server
```bash
go run main.go
```

## API Documentation
Access Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```

## API Routes

### Auth
- POST `/api/v1/auth/register` - Register new user
- POST `/api/v1/auth/login` - Login user
- POST `/api/v1/auth/name` - Edit user name
- POST `/api/v1/auth/email` - Edit user email  
- POST `/api/v1/auth/password` - Edit user password

### Notes
- POST `/api/v1/notes` - Create note
- PUT `/api/v1/notes/{id}` - Edit note
- DELETE `/api/v1/notes/{id}` - Delete note
- GET `/api/v1/notes/{id}` - Get note by ID
- GET `/api/v1/notes` - Get all notes

## Testing
Use the provided `request.rest` file with REST Client plugin to test the APIs.