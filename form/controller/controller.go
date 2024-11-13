package controller

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"formbuilder/connection"
	"formbuilder/models"
	"log"
	"net/smtp"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

	return c.Status(fiber.StatusCreated).JSON(user)
}

// func ForgetPassword(c *fiber.Ctx) error {
// 	type Email struct {
// 		Email string
// 	}
// 	userEmail := new(Email)
// 	if err := c.BodyParser(userEmail); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Cannot parse JSON",
// 		})
// 	}
// 	log.Println(userEmail.Email, "emaillllllll")
// 	var storedUser models.User
// 	err := connection.UserCollection.FindOne(context.TODO(), bson.M{"email": userEmail.Email}).Decode(&storedUser)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error": "Invalid email or password",
// 		})
// 	}
// 	token, err := GenerateToken()
// 	go sendEmail(token, "dhananjay.social26@gmail.com")
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"hello": "helloo"})
// }

func ForgetPassword(c *fiber.Ctx) error {
	type Email struct {
		Email string `json:"email"`
	}
	userEmail := new(Email)
	if err := c.BodyParser(userEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	log.Println(userEmail.Email, "email received")
	var storedUser models.User
	err := connection.UserCollection.FindOne(context.TODO(), bson.M{"email": userEmail.Email}).Decode(&storedUser)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email address",
		})
	}

	// Generate the token and expiration time
	token, err := generateToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate reset token",
		})
	}

	// Set token expiration time (e.g., 1 hour from now)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Update user record with reset token and expiration time
	update := bson.M{
		"$set": bson.M{
			"reset_token":   token,
			"reset_expires": expirationTime,
		},
	}
	_, err = connection.UserCollection.UpdateOne(context.TODO(), bson.M{"email": userEmail.Email}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save reset token",
		})
	}

	// Send email with the reset link
	resetURL := fmt.Sprintf("https://yourdomain.com/reset-password?token=%s", token)
	go sendEmail(resetURL, userEmail.Email)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Password reset email sent"})
}

func ChangeForgetPassword(c *fiber.Ctx) error {
	tokenid := c.Params("id")
	type usernewpass struct {
		Password string
	}
	usernew := new(usernewpass)
	if err := c.BodyParser(usernew); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	type forgetUser struct {
		Password     string    `bson:"password"`
		ResetExpires time.Time `bson:"reset_expires"`
		UpdatedTime  time.Time `bson:"updated_time"`
		ResetToken   string    `bson:"reset_token"`
	}
	forgetUserData := new(forgetUser)
	err := connection.UserCollection.FindOne(context.TODO(), bson.M{"reset_token": tokenid}).Decode(&forgetUserData)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token is not valid ",
		})
	}
	log.Println(forgetUserData)
	if time.Now().After(forgetUserData.ResetExpires) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token has expired",
		})
	}
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(usernew.Password), bcrypt.DefaultCost)
	update := bson.M{
		"$set": bson.M{
			"password":     string(newHashedPassword),
			"updated_time": time.Now(),
		},
		"$unset": bson.M{
			"reset_token":   "",
			"reset_expires": "",
		},
	}
	_, err = connection.UserCollection.UpdateOne(context.TODO(), bson.M{"reset_token": tokenid}, update)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save reset token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Password reset successfully"})
}

func ChangePassword(c *fiber.Ctx) error {
	type userPassword struct {
		Email       string
		NewPassword string
		OldPassword string
	}
	userPass := new(userPassword)

	if err := c.BodyParser(userPass); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	// newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPass.Email), bcrypt.DefaultCost)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "hash password ",
	// 	})
	// }
	// oldhash := string(newHashedPassword)
	log.Println(userPass, "dhfiho")

	// type newUserPassword struct {
	// 	Password    string    `bson:"password"`
	// 	UpdatedTime time.Time `bson:"updated_time"`
	// }

	// newuserPass := new(newUserPassword)

	// err = connection.UserCollection.FindOne(context.TODO(), bson.M{"password": oldhash}).Decode(&newuserPass)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "connection error",
	// 	})
	// }
	// log.Println(newuserPass, "ifi")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "password changed sucesspully"})
}

func sendEmail(token string, mail string) {
	// SMTP server configuration.
	smtpHost := "smtp.gmail.com" // Replace with your SMTP host (e.g., smtp.gmail.com for Gmail)
	smtpPort := "587"            // Port for TLS/STARTTLS, use 465 for SSL

	// Sender data.
	from := "dhananjaykachure.dev@gmail.com"
	password := "qxfc tdla vfvp fayx"

	// Receiver email address.
	to := []string{mail}

	// Message.
	subject := "Subject: Forget password\n"
	url := "https://example.com/verification?token=" + token
	body := "This is a forget password link  from Go!:\n" + url
	message := []byte(subject + "\n" + body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	fmt.Println("Email sent successfully!")
}

func generateToken() (string, error) {
	// Create a byte slice to hold random bytes
	bytes := make([]byte, 16) // Adjust size for longer or shorter tokens if needed

	// Generate random bytes
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convert bytes to a hex string
	return hex.EncodeToString(bytes), nil
}
