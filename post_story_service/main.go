package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type PostStory struct {
	gorm.Model
	Body string    `json:"body" gorm:"type:varchar(1000)"`
	Date time.Time `json:"date"`
}

func setupDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("post_story.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&PostStory{})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	setupDB()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("<h1>hello world</h1>")
	})
	app.Get("/api/post-stories", indexPost)
	app.Get("/api/post-stories/:id", getPost)
	app.Post("/api/post-stories/", createPost)
	app.Patch("/api/post-stories/:id", updatePost)
	app.Delete("/api/post-stories/:id", deletePost)
	log.Fatal(app.Listen(":8084"))

}

// GET POST
func getPost(c *fiber.Ctx) error {
	var post PostStory

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid ID", "Error": err.Error()})
	}

	if err := db.Find(&post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "post not found", "Error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "get post function ", "Post": post})
}

// INDEX POSTS
func indexPost(c *fiber.Ctx) error {
	var posts []PostStory
	if err := db.Find(&posts).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "not found error", "Error ": err.Error()})
	}
	return c.JSON(posts)
}

// CREATE POST
func createPost(c *fiber.Ctx) error {
	var post PostStory
	if err := c.BodyParser(&post); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "invalid error server", "Error": err.Error()})
	}
	post.Date = time.Now()
	db.Create(&post)
	return c.JSON(post)
}

// UPDATE POST
func updatePost(c *fiber.Ctx) error {
	var post PostStory
	var updatedpost PostStory
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid ID", "Error": err.Error()})
	}
	if err := db.Find(&post, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "post not found", "Error": err.Error()})
	}
	if err := c.BodyParser(&updatedpost); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid body", "Error": err.Error()})
	}
	post.Body = updatedpost.Body
	post.Date = updatedpost.Date
	if err := db.Save(&post).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "server error", "Error": err.Error()})
	}
	return c.JSON(post)

}

// DELETE POST
func deletePost(c *fiber.Ctx) error {
	var post PostStory
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid ID", "Error": err.Error()})
	}
	if err := db.Find(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "post not found", "Error": err.Error()})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "server error", "Error": err.Error()})
	}
	if err := db.Delete(&post).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "server error", "Error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "successfully deleted"})
}
