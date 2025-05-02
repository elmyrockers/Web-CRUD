package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/jinzhu/copier"

	"Web_CRUD/model"
	"Web_CRUD/dto"

	"log"
	// "fmt"
	"github.com/davecgh/go-spew/spew"
)


func NewUserRepository(db *sqlx.DB) *UserRepository {
	base := NewBaseRepository( db )
	return &UserRepository{ base: base }
}

type UserRepository struct{
	base *BaseRepository
}

func (u *UserRepository) All() []model.User {
	// List of all rows
		var users []model.User
		err := u.base.DB.Select(&users, "SELECT * FROM users")
		if err != nil {
			log.Fatalln(err)
		}

	return users
}

func (u *UserRepository) GetByID(id string ) model.User {
	// Get single user by ID
		var user model.User
		err := u.base.DB.Get(&user, "SELECT * FROM users WHERE id=?", id )
		if err != nil {
			log.Fatalln(err)
		}

	return user
}

func (u *UserRepository) New(request dto.NewUserRequest ) ( int64, error) {
	// Map request to db model
		var user model.User
		err := copier.Copy( &user, &request )
		if err != nil {
			log.Fatal(err)
		}
	// Create new user record
		result, err := u.base.DB.NamedExec( "INSERT INTO users(name, email, password) VALUES(:name, :email, :password)", user )
		if err != nil {
			return 0, err
		}
		userID, _ := result.LastInsertId()

		spew.Dump( user, userID, err )

	// return userID
		return userID, err
}

func (u *UserRepository) Create() {
	
}