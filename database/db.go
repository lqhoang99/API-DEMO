package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/samuel/go-zookeeper/zk"
)

var DB *mongo.Database
var Zoo *zk.Conn

//ConnectZookeeper ...
func ConnectZookeeper() {
	c, _, err := zk.Connect([]string{"zookeeper:2181"}, time.Second) 
	if err != nil {
		panic(err)
	}
	Zoo=c
}
//GetValueFromZoo ...
func GetValueFromZoo(path string) string{
	res,_,_:=Zoo.Get(path)
	return string(res)
}
//Connectdb ...
func Connectdb(dbName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://db:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	DB = client.Database(dbName)

	fmt.Println("Connected to db:", dbName)
	return DB
}
