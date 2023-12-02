package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	Title       string    `json:"title" gorm:"type:varchar(250)"`
	Description string    `json:"description" gorm:"type:varchar(500)"`
	Body        string    `json:"body" gorm:"type:varchar(1000)"`
	Photopath   string    `json:"photopath" gorm:"type:varchar(250)"`
	Date        time.Time `json:"date"`
}

func setupDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("news-story.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&Post{})
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
	app.Get("/api/post-news", indexPost)
	app.Get("/api/post-news/:id", getPost)
	app.Post("/api/post-news/", createPost)
	app.Patch("/api/post-news/:id", updatePost)
	app.Delete("/api/post-news/:id", deletePost)
	log.Fatal(app.Listen(":8083"))

}

// GET POST
func getPost(c *fiber.Ctx) error {
	var post Post

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
	var posts []Post
	if err := db.Find(&posts).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "not found error", "Error ": err.Error()})
	}
	return c.JSON(posts)
}

// CREATE POST
func createPost(c *fiber.Ctx) error {
	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Не удалось получить файл",
		})
	}

	// Генерируем уникальное имя файла
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
	// Путь к папке uploads
	uploadPath := "./uploads"
	// Полный путь к файлу
	filePath := filepath.Join(uploadPath, fileName)

	// Сохраняем файл в папку uploads
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось сохранить файл",
		})
	}

	var post Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "invalid error server", "Error": err.Error()})
	}
	post.Date = time.Now()
	post.Photopath = filePath
	if err := db.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось сохранить пост",
		})
	}

	return c.JSON(post)
}

// UPDATE POST
func updatePost(c *fiber.Ctx) error {
	var post Post
	var updatedPost Post

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Неверный ID", "Error": err.Error()})
	}

	// Находим пост по ID
	if err := db.Find(&post, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Пост не найден", "Error": err.Error()})
	}

	// Парсим обновленные данные из запроса
	if err := c.BodyParser(&updatedPost); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Неверное тело запроса", "Error": err.Error()})
	}

	// Обновляем данные поста
	post.Body = updatedPost.Body
	post.Date = updatedPost.Date

	// Проверяем, был ли передан новый файл
	newFile, newFileErr := c.FormFile("photo")
	if newFileErr == nil && newFile != nil {
		// Генерируем уникальное имя файла
		fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(newFile.Filename))
		// Путь к папке uploads
		uploadPath := "./uploads"
		// Полный путь к новому файлу
		filePath := filepath.Join(uploadPath, fileName)

		// Сохраняем новый файл в папку uploads
		if err := c.SaveFile(newFile, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Не удалось сохранить новый файл",
			})
		}

		// Обновляем путь к фотографии в посте
		post.Photopath = filePath
	}

	// Сохраняем обновленные данные в базе данных
	if err := db.Save(&post).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка сервера", "Error": err.Error()})
	}

	return c.JSON(post)
}

// DELETE POST
func deletePost(c *fiber.Ctx) error {
	var post Post
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
