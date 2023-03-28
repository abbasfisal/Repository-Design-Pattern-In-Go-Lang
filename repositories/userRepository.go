package repositories

import (
	"database/sql"
	"tut-youtube/models"
)

//work on db
type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (userRepo UserRepository) Create(u models.User) (models.User, error) {
	//you can use userRepo.Db to working with mysql db
	//userRepo.Db.Exec()
	return u, nil
}
