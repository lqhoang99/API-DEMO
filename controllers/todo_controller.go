package controllers

import (
	"RestAPI/database"
	"RestAPI/models"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateTodo func
func CreateTodo(c echo.Context) error {
	db := database.Connectdb()
	collection := db.Collection("todos")
	todo := new(models.Todo)
	c.Bind(todo)
	log.Println(todo)

	res, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(res)
	return c.JSON(http.StatusCreated, res.InsertedID)
}


