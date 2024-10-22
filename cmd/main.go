package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_userController "github.com/eskimoburger/cursor-pagination/controllers"
	"github.com/eskimoburger/cursor-pagination/entities"
	_userRepository "github.com/eskimoburger/cursor-pagination/repositories"
	_userService "github.com/eskimoburger/cursor-pagination/services"
)

func main() {
	users := []entities.User{
		{ID: 1, Name: "Alice", CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
		{ID: 2, Name: "Bob", CreatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)},
		{ID: 3, Name: "Charlie", CreatedAt: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC)},
		{ID: 4, Name: "David", CreatedAt: time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC)},
		{ID: 5, Name: "Eve", CreatedAt: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC)},
		{ID: 6, Name: "Frank", CreatedAt: time.Date(2023, 1, 6, 0, 0, 0, 0, time.UTC)},
		{ID: 7, Name: "Grace", CreatedAt: time.Date(2023, 1, 7, 0, 0, 0, 0, time.UTC)},
		{ID: 8, Name: "Heidi", CreatedAt: time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC)},
		{ID: 9, Name: "Ivan", CreatedAt: time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC)},
		{ID: 10, Name: "Judy", CreatedAt: time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)},
		{ID: 11, Name: "Mallory", CreatedAt: time.Date(2023, 1, 11, 0, 0, 0, 0, time.UTC)},
		{ID: 12, Name: "Niaj", CreatedAt: time.Date(2023, 1, 12, 0, 0, 0, 0, time.UTC)},
		{ID: 13, Name: "Olivia", CreatedAt: time.Date(2023, 1, 13, 0, 0, 0, 0, time.UTC)},
		{ID: 14, Name: "Peggy", CreatedAt: time.Date(2023, 1, 14, 0, 0, 0, 0, time.UTC)},
		{ID: 15, Name: "Quentin", CreatedAt: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)},
	}
	userRepository := _userRepository.NewInMemoryUserRepository(users)
	userService := _userService.NewUserService(userRepository)
	userHandler := _userController.NewUserController(userService)

	http.Handle("/users", userHandler)
	fmt.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
