package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/template/jet/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"Web_CRUD/controller"

	// "github.com/davecgh/go-spew/spew"
	// "fmt"
	"os"
	"log"
	"time"
)

func createSessionStore() *session.Store {
	return session.New(session.Config{
		Expiration: 30 * time.Minute, // Auto-delete inactive sessions
	})
}

func connectDB() *sqlx.DB {
	// Get database config
		dsn := os.Getenv( "DSN" )
		// dsn := "root@tcp(127.0.0.1:3306)/golangdb?parseTime=true"
		// fmt.Println( "test sahaja: ", dsn )
		// fmt.Println("DSN:", os.Getenv("DSN"))

	// Connect to MySQL database
		db, err := sqlx.Connect( "mysql", dsn )
		if err != nil {
			log.Fatalln( "Failed to connect to DB:", err )
		}

	// Set up connection pool options
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(5 * time.Minute)

	return db
}

func main() {
	// Initialize Fiber with the Jet template engine
		engine := jet.New("./view", ".jet")
		app := fiber.New(fiber.Config{
			Views: engine,
		})

	// Initialize controllers
		db := connectDB()
		store := createSessionStore()
		userController := controller.NewUserController(db,store)

	// Routing
		app.Static("/", "./public")

		// Route for Home
			app.Get("/",func(c *fiber.Ctx) error {
				return c.SendString( "This is home page" )
			})

		// Routes for User
			app.Get("/users", userController.List)
			app.Get("/users/new", userController.New)
			app.Get("/users/edit/:id", userController.Edit)
			app.Get("/users/delete/:id", userController.Delete)

			app.Post("/users/new", userController.New_process)
			app.Post("/users/edit/:id", userController.Edit_process)
			app.Post("/users/delete/:id", userController.Delete_process)

	app.Listen(":3000")
}