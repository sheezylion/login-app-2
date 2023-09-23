package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Define a User model for the database
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-"`
}

var db *gorm.DB

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Construct the database DSN from environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Initialize the database connection
	var errDb error
	db, errDb = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDb != nil {
		log.Fatal(errDb)
	}

	// Auto Migrate the User model to create the users table
	db.AutoMigrate(&User{})

	// Initialize the Gin router
	r := gin.Default()

	// Define API routes
	api := r.Group("/api")
	{
		api.POST("/register", register)
		api.POST("/login", login)
		api.GET("/health", HealthCheckHandler)
	}

	// Start the server
	port := "4000"
	r.Run(":" + port)
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

// HealthCheckHandler is the handler function for the health check endpoint.
func HealthCheckHandler(c *gin.Context) {
	// You can perform any health checks here. For simplicity, we'll just return a success status.
	response := HealthCheckResponse{Status: "OK"}
	c.JSON(http.StatusOK, response)
}

func register(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	// Check if the username already exists
	var existingUser User
	result := db.Where("username = ?", user.Username).First(&existingUser)
	if result.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username already exists"})
		return
	}

	// Hash the user's password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	user.Password = string(hashedPassword)

	// Create the user in the database
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func login(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	// Query the database for all users with the given username
	var users []User
	result := db.Where("username = ?", user.Username).Find(&users)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// Loop through the fetched users and compare passwords
	for _, u := range users {
		if u.Password == user.Password {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
			return
		}
	}

	// No matching user with correct password found
	c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})

}
