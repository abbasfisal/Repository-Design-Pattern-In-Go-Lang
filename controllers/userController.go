package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"tut-youtube/models"
	"tut-youtube/repositories"
)

//repository interface
type BaseHandler struct {
	UserRepo repositories.UserRepositoryInterface
}

func NewBaseHandler(userRepo repositories.UserRepositoryInterface) *BaseHandler {
	return &BaseHandler{UserRepo: userRepo}
}

func (h *BaseHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	u := models.User{
		Name:  "ali",
		Phone: "09232923",
	}
	user, err := h.UserRepo.Create(u)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(user)
}
