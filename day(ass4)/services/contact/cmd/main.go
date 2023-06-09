package main

import (
	"Projects/day/pkg/store/postgres"
	"Projects/day/services/contact/internal/delivery"
	"Projects/day/services/contact/internal/repository"
	"Projects/day/services/contact/internal/useCase"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := postgres.OpenDB("localhost", "", "postgres", "1234", "day")
	if err != nil {
		fmt.Printf("Fail connecting ")
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userUseCase := useCase.NewUserUseCase(userRepository)
	userDelivery := delivery.NewUserDelivery(userUseCase)

	http.HandleFunc("/users", userDelivery.HandleUsersRequest)
	http.HandleFunc("/users/{id}", userDelivery.HandleUserRequest)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error server starting  ", err)
	}
}
