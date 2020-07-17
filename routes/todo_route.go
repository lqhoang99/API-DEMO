package routes

import (
	"RestAPI/controllers"
	"github.com/labstack/echo"
)
//TodoRoute func
func TodoRoute(g *echo.Group) {
	//Creat
	g.POST("/Create", controllers.CreateTodo)
	
}
