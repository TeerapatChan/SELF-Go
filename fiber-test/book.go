package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Book
// @Router /books [get]
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Book not found")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)

	if c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == id {
			book.Title = bookUpdate.Title
			book.Author = bookUpdate.Author
			books[i] = book
			return c.JSON(book)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Book not found")
}

func deleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return c.JSON(book)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Book not found")
}
