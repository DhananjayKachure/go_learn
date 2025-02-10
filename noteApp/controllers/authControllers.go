package controllers

import (
	"context"
	"log"
	"noteapp/config"
	"noteapp/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()
	collection := config.GetCollection("users")
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to register user"})
	}
	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	collection := config.GetCollection("users")

	err := collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Credentials"})
	}

	// Ensure JWT_SECRET is set
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("JWT_SECRET is empty or not set")
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Create JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID.Hex(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign Token
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("JWT signing error:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
	}

	log.Println("Generated Token:", tokenStr) // Debugging

	return c.JSON(fiber.Map{"token": tokenStr})
}
