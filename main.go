package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	app := fiber.New()

	// Mount different services
	app.Mount("/post-stories/", PostStoryService())
	app.Mount("/post-news/", PostNewsService())
	app.Mount("/staff/", StaffService())
	app.Mount("/staff-position/", StaffPositionService())
	app.Mount("/pride-student/", PrideStudentService())

	// Run the application on port 3000
	app.Listen(":3000")
}

func PrideStudentService() http.Handler {
	return http.StripPrefix("/pride-student/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "http://localhost:8000"+r.URL.Path)
	}))
}

func StaffPositionService() http.Handler {
	// Implement the logic for StaffPositionService
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your implementation here
	})
}

func StaffService() http.Handler {
	// Implement the logic for StaffService
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your implementation here
	})
}

func PostNewsService() http.Handler {
	// Implement the logic for PostNewsService
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your implementation here
	})
}

func PostStoryService() http.Handler {
	return http.StripPrefix("/post-stories/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "http://localhost:8081"+r.URL.Path)
	}))
}
