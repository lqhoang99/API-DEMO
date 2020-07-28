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
var idCompleted = primitive.NewObjectID()
var idUpdate = primitive.NewObjectID()
var idDelete = primitive.NewObjectID()
func (s TodoSuite) SetupSuite() {
	//url:=database.GetValueFromZoo("/mongodb")
	database.Connectdb("todo-test")
	removeOldData()
	addRecord(idCompleted) // for test Completed
	addRecord(idUpdate) // for test Update
	addRecord(idDelete) // for test Delete

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
	assert.Equal(s.T(), res.Completed, todo.Completed)

}

func ToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

// Test Completed

func addRecord(id primitive.ObjectID) {
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
	c.SetParamValues(idCompleted.Hex())

	controllers.Complete(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	res:= struct {
		MatchedCount int 			`bson:"MatchedCount" json:"MatchedCount"`
		ModifiedCount int			`bson:"ModifiedCount" json:"ModifiedCount"`
		UpsertedCount int           `bson:"UpsertedCount" json:"UpsertedCount"`
		UpsertedID interface{}      `bson:"UpsertedID" json:"UpsertedID"`
	}{

	}
	json.Unmarshal([]byte(rec.Body.String()), &res)

	assert.Equal(s.T(), res.MatchedCount, 1)
	assert.Equal(s.T(), res.ModifiedCount, 1)
	assert.Equal(s.T(), res.UpsertedCount, 0)
	assert.Equal(s.T(), res.UpsertedID, nil)
}

// Test GetList

func (s *TodoSuite) TestGetList() {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	controllers.GetList(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)

	ctx := context.Background()
	cursor, err := database.DB.Collection("todos").Find(ctx, bson.M{})
	if err != nil {
		panic("query err")
	}

	var todos []models.Todo
	defer cursor.Close(ctx)
	cursor.All(ctx, &todos)

	var res []models.Todo
	json.Unmarshal(rec.Body.Bytes(), &res)
	assert.Equal(s.T(), todos, res)

}
//Test Update

func (s *TodoSuite) TestUpdate() {

	e := echo.New()
	todo := models.Todo{
		Title: "hoangdeptrai",
		Desc:  "qua dep trai",
		Completed:false,
	}
	req := httptest.NewRequest(http.MethodPut, "/todos/:id", ToIOReader(todo))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(idUpdate.Hex())

	controllers.Update(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	res:= struct {
		MatchedCount int 			`bson:"MatchedCount" json:"MatchedCount"`
		ModifiedCount int			`bson:"ModifiedCount" json:"ModifiedCount"`
		UpsertedCount int           `bson:"UpsertedCount" json:"UpsertedCount"`
		UpsertedID interface{}      `bson:"UpsertedID" json:"UpsertedID"`
	}{

	}
	json.Unmarshal([]byte(rec.Body.String()), &res)

	assert.Equal(s.T(), res.MatchedCount, 1)
	assert.Equal(s.T(), res.ModifiedCount, 1)
	assert.Equal(s.T(), res.UpsertedCount, 0)
	assert.Equal(s.T(), res.UpsertedID, nil)

}

//Test Delete
func (s *TodoSuite) TestDelete() {

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/todos/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(idDelete.Hex())

	controllers.Delete(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	x :=struct{
		DeletedCount int	`bson:"DeletedCount" json:"DeletedCount"`
	}{

	}
	json.Unmarshal([]byte(rec.Body.String()), &x)
	assert.Equal(s.T(), x.DeletedCount, 1)

}


func TestTodoSuite(t *testing.T) {
	suite.Run(t, new(TodoSuite))
}
