package routes

import (
	"RestAPI/controllers"

	"github.com/labstack/echo"
)

//TodoRoute func
func TodoRoute(g *echo.Group) {
	// /todos
	g.POST("", controllers.CreateTodo)

	// todos/:id/completed
	g.PATCH("/:id/completed", controllers.Complete)

	// /todos
	g.GET("", controllers.GetList)

	//
	g.DELETE("/:id", controllers.Delete)

	g.PUT("/:id", controllers.Update)
}
