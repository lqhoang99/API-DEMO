package main

import (
	"RestAPI/database"
	"RestAPI/routes"

	"github.com/labstack/echo"
)

func init() {
	database.Connectdb("todos")
}

func main() {
	server := echo.New()
	routes.TodoRoute(server.Group("/todos"))
	server.Logger.Fatal(server.Start(":8080"))
}
