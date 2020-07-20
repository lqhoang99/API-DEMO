package test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"

	"RestAPI/controllers"
	"RestAPI/database"
	"RestAPI/models"
)

type CreateTodoSuite struct {
	suite.Suite
	Todos []models.Todo
}

func (s CreateTodoSuite) SetupSuite()  {
	database.Connectdb("todo-test")

	removeOldData()
}

func (s CreateTodoSuite) TearDownSuite() {
	removeOldData()
}

func removeOldData()  {
	database.DB.Collection("todos").DeleteMany(context.Background(), bson.M{})
}

func (s *CreateTodoSuite) TestCreateTodo() {

	e := echo.New()

	todo := models.Todo{
		Title: "title 1",
		Desc: "aab",
	}
	req := httptest.NewRequest(http.MethodPost, "/todos", ToIOReader(todo))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controllers.CreateTodo(c)

	assert.Equal(s.T(), http.StatusCreated, rec.Code)

	var res models.Todo

	json.Unmarshal([]byte(rec.Body.String()), &res)

	assert.Equal(s.T(), res.Title, todo.Title)
	assert.Equal(s.T(), res.Desc, todo.Desc)

	// ctx := context.Background()
	// cursor, err := database.DB.Collection("todos").Find(ctx, bson.M{})
	// if err != nil {
	// 	panic("query err")
	// }
	//
	// var todos []models.Todo
	//
	// defer cursor.Close(ctx)
	// cursor.All(ctx, &todos)
	//
	// assert.Equal(s.T(), len(todos), 1)
}

func ToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

func TestCreateTodoSuite(t *testing.T)  {
	suite.Run(t, new(CreateTodoSuite))
}

