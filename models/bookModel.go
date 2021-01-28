package models

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Book model
type Book struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:binary(16);primary_key;default:(UUID_TO_BIN(UUID()));"`
	Title     string    `json:"name"`
	Author    string    `json:"author"`
	Price     string    `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// BeforeCreate method
func (b *Book) BeforeCreate(scope *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

// GetBooks function
func GetBooks(c *fiber.Ctx) error {
	var books []Book
	if db.Find(&books).Error != nil {
		return &fiber.Error{}
	}
	return c.JSON(fiber.Map{"data": books})
}

// GetBook function
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book
	if db.Find(&book, id).Error != nil {
		return &fiber.Error{}
	}
	return c.JSON(book)
}

// NewBook func
func NewBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString("Parse error")
	}
	db.Create(&book)
	return c.JSON(book)
}

// DeleteBook func
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No Book Found with ID")
	}
	db.Delete(&book)
	return c.SendString("Book Successfully deleted")
}
