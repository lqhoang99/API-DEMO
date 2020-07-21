package controllers

import (
	"RestAPI/database"
	"RestAPI/models"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateTodo func
func CreateTodo(c echo.Context) error {
	collection := database.DB.Collection("todos")
	todo := new(models.Todo)
	c.Bind(todo)
	todo.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusCreated, todo)
}

//Complete func
func Complete(c echo.Context) error {
	collection := database.DB.Collection("todos")
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, updateResult)
}

//GetList func
func GetList(c echo.Context) error {
	collection := database.DB.Collection("todos")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var results []*models.Todo
	for cursor.Next(context.TODO()) {
		var elem models.Todo
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	cursor.Close(context.TODO())
	return c.JSON(http.StatusOK, results)
}

// Delete func
func Delete(c echo.Context) error {
	collection := database.DB.Collection("todos")
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, deleteResult)
}

//Update func
func Update(c echo.Context) error {
	collection := database.DB.Collection("todos")
	id := c.Param("id")
	todo := new(models.Todo)
	c.Bind(todo)

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"title":     todo.Title,
		"desc":      todo.Desc,
		"completed": todo.Completed,
	}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, updateResult)
}
