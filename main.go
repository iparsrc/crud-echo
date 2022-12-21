package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	// create a new repository
	repository := NewRepository()

	// create a new server
	server := NewServer(repository)

	// create a new router
	router := echo.New()

	// map routes
	router.GET("/users/:email", server.GetUser)
	router.POST("/users", server.CreateUser)
	router.PUT("/users/:email", server.UpdateUser)
	router.DELETE("/users/:email", server.DeleteUser)

	// start router
	router.Start(":8000")
}
