package main

import (
	"fmt"
	"log"
	"os"
	"replog/internal/handlers"
	"replog/internal/repository"
	"replog/internal/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var db *sqlx.DB

	for i := 0; i < 10; i++ {
		db, err = sqlx.Connect("postgres", DSN)
		if err == nil {
			break
		}

		log.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	fmt.Println("database connected successfully")
	fmt.Println("database connected successfully")

	//dependency injection
	UserRepo := repository.NewUserRepository(db)
	UserService := services.NewAuthService(UserRepo)
	UserHandler := handlers.NewAuthHandler(UserService)

	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", UserHandler.RegisterUser)
	}

	fmt.Println("server is starting at port 8080....")
	r.Run(":8080")

}
