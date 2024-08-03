package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/TeerapatChan/fiber-test/docs"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/swagger"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var books []Book

var myUser = User{
	Email:    "user@example.com",
	Password: "password",
}

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Get("/swagger/*", swagger.HandlerDefault)
	books = append(books, Book{ID: 1, Title: "Book 1", Author: "Author 1"})
	books = append(books, Book{ID: 2, Title: "Book 2", Author: "Author 2"})
	books = append(books, Book{ID: 3, Title: "Book 3", Author: "Author 3"})

	app.Post("/login", login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	app.Use(checkMiddleware)

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/config", getEnv)

	app.Listen(":8080")
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File uploaded")
}

func getEnv(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}

func checkMiddleware(c *fiber.Ctx) error {
	fmt.Println("Time:", time.Now())
	fmt.Printf("Request: %s\n", c.OriginalURL())
	fmt.Printf("Method: %s\n", c.Method())
	fmt.Printf("IP: %s\n", c.IP())

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	return c.Next()
}

func login(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != myUser.Email || user.Password != myUser.Password {
		return fiber.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["role"] = "admin" // example role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	secretKey := os.Getenv("JWT_SECRET")
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   t,
	})
}
