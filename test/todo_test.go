package test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"

	"RestAPI/controllers"
	"RestAPI/database"
	"RestAPI/models"
)

type TodoSuite struct {
	suite.Suite
	Todos []models.Todo
}

func (s TodoSuite) SetupSuite() {
	database.Connectdb("todo-test")
	removeOldData()
	addRecord()// for test Completed
}

func (s TodoSuite) TearDownSuite() {
	//removeOldData()
}

func removeOldData() {
	database.DB.Collection("todos").DeleteMany(context.Background(), bson.M{})
}

//Test TestCreateTodo
func (s *TodoSuite) TestCreateTodo() {

	e := echo.New()

	todo := models.Todo{
		Title: "title 1",
		Desc:  "aab",
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


}

func ToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

// Test Completed
var id = primitive.NewObjectID()
func addRecord() {
	todo := models.Todo{
		ID:        id,
		Title:     "title 2",
		Desc:      "desc 2",
		Completed: false,
	}
	database.DB.Collection("todos").InsertOne(context.TODO(), todo)
}
func (s *TodoSuite) TestComplete() {

	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/todos/:id/completed", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id.Hex())

	controllers.Complete(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
 	var res models.ResBody
	json.Unmarshal([]byte(rec.Body.String()), &res)

	assert.Equal(s.T(), res.MatchedCount,1)
	assert.Equal(s.T(), res.ModifiedCount,1)
	assert.Equal(s.T(), res.UpsertedCount,0)
	assert.Equal(s.T(), res.UpsertedID,nil)
}

// Test 

func TestTodoSuite(t *testing.T) {
	suite.Run(t, new(TodoSuite))
}
