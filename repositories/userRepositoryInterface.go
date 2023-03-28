package repositories

import "tut-youtube/models"

type UserRepositoryInterface interface {
	Create(u models.User) (models.User, error)
}
