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
	"gorm.io/driver/mysql"
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

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Obtener las credenciales de la base de datos desde las variables de entorno
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Crear la cadena de conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Conectar a la base de datos
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
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
		log.Fatalf("Error loading .env fil with db config: %v", err)
	}
}
