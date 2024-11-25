package controller

import (
	"context"
	"formbuilder/connection"
	"formbuilder/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	_, err = connection.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func LoginUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	var storedUser models.User
	err := connection.UserCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&storedUser)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(storedUser)
}

func ForgetPassword(c *fiber.Ctx) error {
	type Email struct {
		Email string
	}
	userEmail := new(Email)
	if err := c.BodyParser(userEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	log.Println(userEmail.Email, "emaillllllll")
	var storedUser models.User
	err := connection.UserCollection.FindOne(context.TODO(), bson.M{"email": userEmail.Email}).Decode(&storedUser)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	log.Println(storedUser.Email, "userr")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"hello": "helloo"})
}

func GetUser(c *fiber.Ctx) error {
	// Slice to store the raw user data
	var users []bson.M

	// Fetch all documents from the UserCollection
	cursor, err := connection.UserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	defer cursor.Close(context.TODO())

	// Iterate through the cursor and decode each document into the slice
	if err := cursor.All(context.TODO(), &users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode user data",
		})
	}

	// Check if no users were found
	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No users found",
		})
	}

	// Return the raw user data as JSON
	return c.Status(fiber.StatusOK).JSON(users)
}

func GetsectionListing(c *fiber.Ctx) error {
	// Query parameters for pagination
	page := c.QueryInt("page", 1) // Default to page 1 if not provided
	if page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page number",
		})
	}

	limit := c.QueryInt("limit", 10) // Default to 10 items per page if not provided
	if limit < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid limit number",
		})
	}

	// Calculate the skip value
	skip := (page - 1) * limit

	// Fetch paginated data from the UserCollection
	cursor, err := connection.UserCollection.Find(context.TODO(), bson.M{},
		options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	defer cursor.Close(context.TODO())

	// Decode the documents into a slice
	var topics []bson.M
	if err := cursor.All(context.TODO(), &topics); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode user data",
		})
	}

	// Check if no users were found
	if len(topics) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No users found",
		})
	}

	// Count total documents for pagination metadata
	totalCount, err := connection.UserCollection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to count users",
		})
	}

	// Return paginated response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"page":   page,
		"limit":  limit,
		"total":  totalCount,
		"topics": topics,
	})
}
