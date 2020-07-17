package main

import (
	"RestAPI/routes"

	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	routes.TodoRoute(server.Group("/todos"))
	server.Logger.Fatal(server.Start(":8080"))
}
