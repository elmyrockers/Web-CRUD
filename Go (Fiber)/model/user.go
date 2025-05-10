package model

import "time"

type User struct {
	ID		  int		`db:"id"`
	Name	  string	`db:"name"`
	Email	  string	`db:"email"`
	Website   string	`db:"website"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}