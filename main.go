package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/controllers"
	"github.com/heroku/go-getting-started/middleware"
	"github.com/heroku/go-getting-started/repository"
	"github.com/heroku/go-getting-started/utils"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	host := os.Getenv("DB_HOST")
	DBport := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, DBport, user, dbname, pass)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connectin to DB: %v", err)
	}

	repo := repository.NewUserRepo(db)
	tokengen := utils.JWTTokenService{}
	userController := controllers.NewUserController(repo, tokengen)
	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", userController.Register)
		userGroup.POST("/login", userController.Login)
		userGroup.PUT("/update", middleware.Authenticate(), userController.UpdateUser)
	}

	router.GET("/healt", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "All OK, service is running.",
		})
	})

	router.Run(":" + port)
}

func init() {
	// Carga el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
