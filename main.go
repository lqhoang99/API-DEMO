package main

import (
	"RestAPI/database"
	"RestAPI/routes"
	"github.com/labstack/echo"
)

func init() {
	database.ConnectZookeeper()
	//url:=database.GetValueFromZoo("/mongodb")
	//database.Connectdb(url,"todos-list")
	database.Connectdb("todos-list")
}

func main() {
	server := echo.New()
	routes.TodoRoute(server.Group("/todos"))
	server.Logger.Fatal(server.Start(":3000"))
}
