package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/gofiber/fiber/v2/middleware/session"
	// "github.com/jinzhu/copier"

	"Web_CRUD/dto"
	// "Web_CRUD/model"
	"Web_CRUD/repository"
	
	

	"github.com/davecgh/go-spew/spew"
	// "fmt"
	// "log"
)

func NewUserController(db *sqlx.DB, store *session.Store) *UserController {
	userRepository := repository.NewUserRepository(db)
	return &UserController{ UserRepository: userRepository, sessionStore: store }
}








type UserController struct{
	UserRepository *repository.UserRepository
	sessionStore *session.Store
}

func (u *UserController) List(c *fiber.Ctx) error {
	session, err := u.sessionStore.Get( c )
	if err != nil { return err } // Handling for session lost

	flash := session.Get( "flash_message" )
	if flash == nil { // Set default message
		flash = ""
	}
	session.Delete( "flash_message" )
	session.Save()

	users := u.UserRepository.All()
	
	// Display view
		return c.Render( "User/list", fiber.Map{
			"users": users,
			"flash": flash,
		})
}

func (u *UserController) New(c *fiber.Ctx) error {
	// Display view
		return c.Render( "User/new", fiber.Map{} )
}

func (u *UserController) Edit(c *fiber.Ctx) error {
	id := c.Params( "id" )
	user := u.UserRepository.GetByID( id )
	// spew.Dump( user.ID )

	// Display view
		return c.Render( "User/edit", fiber.Map{
			"user": user,
		})
}

func (u *UserController) Delete(c *fiber.Ctx) error {
	id := c.Params( "id" )
	user := u.UserRepository.GetByID( id )
	spew.Dump( user.ID )

	// Display view
		return c.Render( "User/delete", fiber.Map{
			"user": user,
		})
}

func (u *UserController) New_process(c *fiber.Ctx) error {
		session, err := u.sessionStore.Get( c )
		// keys := session.Keys()
		// spew.Dump( keys, err )
		// return nil


	// spew.Dump( store )
	// Get form values
		var request dto.NewUserRequest
		err = c.BodyParser( &request )
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

	// Validate user request
		err = dto.Validate( &request )
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_errors": err.Error()})
		}

	// Save form data into db table
		userID, err := u.UserRepository.New( request )
		spew.Dump( userID )
		if err != nil {
			session.Set( "flash_message", "<div class='alert alert-danger'>Failed to create new user record!</div>" )
			session.Save()
			return c.Redirect( "/users" )
		}
		session.Set( "flash_message", "<div class='alert alert-success'>New user record has been created successfully!</div>" )
		session.Save()
		return c.Redirect( "/users" )
}

func (u *UserController) Edit_process(c *fiber.Ctx) error {
	id := c.Params( "id" )
	user := u.UserRepository.GetByID( id )
	spew.Dump( user.ID )

	// Display view
		return c.Render( "User/delete", fiber.Map{
			"user": user,
		})
}

func (u *UserController) Delete_process(c *fiber.Ctx) error {
	id := c.Params( "id" )
	user := u.UserRepository.GetByID( id )
	spew.Dump( user.ID )

	// Display view
		return c.Render( "User/delete", fiber.Map{
			"user": user,
		})
}