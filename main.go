package main

import (
	"github.com/cesc1802/go_training/internal/route"
	"github.com/cesc1802/go_training/internal/storages"
	"log"
)

func main() {
	db := storages.Connect()
	// create database
	storages.Migration(db)

	testUser := &storages.User{
		ID:       "firstUser",
		Password: "example",
		MaxTodo:  5,
	}

	// check default user exist
	user := &storages.User{}
	db.Where("id = ?", testUser.ID).First(&user)

	if user.ID != "firstUser" {
		userRepo := storages.NewUserRepo()
		userRepo.Save(testUser)
	} else {
		log.Println("Exist")
	}

	route.SetupRoutes(db)
}
