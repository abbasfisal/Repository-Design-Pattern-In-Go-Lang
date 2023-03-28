package main

import (
	"fmt"
	"log"
	"net/http"
	"tut-youtube/controllers"
	"tut-youtube/database"
	"tut-youtube/repositories"
)

func main() {
	//
	db := database.ConnectDb()
	err := db.Ping()

	if err != nil {
		log.Fatal("error in ping db", err)
	}
	fmt.Println("ping success!")

	userRepo := repositories.NewUserRepository(db)
	uController := controllers.NewBaseHandler(userRepo)
	//

	http.HandleFunc("/users/create", uController.CreateUserHandler)

	fmt.Println("start Server on Port : 8080")
	http.ListenAndServe(":8080", nil)
}
