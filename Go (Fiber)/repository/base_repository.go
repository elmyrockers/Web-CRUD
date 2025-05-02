package repository

import (
	"github.com/jmoiron/sqlx"
)

func NewBaseRepository(db *sqlx.DB) *BaseRepository {
	return &BaseRepository{ DB: db }
}

type BaseRepository struct{
	DB *sqlx.DB
}

func (b *BaseRepository) FindAll() {

}