package controllers

import (
	"context"
	"noteapp/config"
	"noteapp/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNote(c *fiber.Ctx) error {
	var notes models.Note
	err := c.BodyParser(&notes)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})

	}
	userIDStr, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "User ID not found in token"})
	}
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	notes.ID = primitive.NewObjectID()
	notes.UserID = userID
	notes.Created = time.Now().Format(time.RFC1123)
	collection := config.GetCollection("notes")
	_, err = collection.InsertOne(context.TODO(), notes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create note"})
	}

	return c.JSON(notes)
}

func GetNotes(c *fiber.Ctx) error {
	userIdStr, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User ID not found in token"})
	}
	UserID, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	filter := bson.M{"user_id": UserID}
	collection := config.GetCollection("notes")
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch notes"})
	}
	defer cursor.Close(context.TODO())
	var notes []models.Note
	if err := cursor.All(context.TODO(), &notes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse notes"})
	}

	return c.JSON(notes)
}
