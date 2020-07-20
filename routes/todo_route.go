package routes

import (
	"RestAPI/controllers"

	"github.com/labstack/echo"
)

//TodoRoute func
func TodoRoute(g *echo.Group) {
	g.POST("/Create", controllers.CreateTodo)
	g.PATCH("/:id/Completed", controllers.Complete)
	g.GET("/List", controllers.GetList)
	g.DELETE("/:id/Delete", controllers.Delete)
	g.PUT("/:id/Update", controllers.Update)
}
