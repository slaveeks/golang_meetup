package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type member struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

func main() {
	e := echo.New()

	ctx := context.TODO()

	collection := connectToDatabase("mongo", "27017", "meetup", "members")

	e.GET("/members", func(c echo.Context) error {
		var m []*member
		cursor, err := collection.Find(ctx, bson.D{})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		err = cursor.All(ctx, &m)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, &m)
	})

	e.POST("/members", func(c echo.Context) error {
		name := c.QueryParam("name")
		email := c.QueryParam("email")
		m := member{
			Name:  name,
			Email: email,
		}

		_, err := collection.InsertOne(ctx, m)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, &m)
	})

	e.DELETE("/members/:id", func(c echo.Context) error {
		id := c.Param("id")
		var m *member

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		filter := bson.M{"_id": oid}

		result := collection.FindOneAndDelete(ctx, filter)

		if result.Err() != nil {
			return c.JSON(http.StatusNotFound, result.Err().Error())
		}

		err = result.Decode(&m)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, &m)
	})

	e.GET("/members/:id", func(c echo.Context) error {
		id := c.Param("id")
		var m *member

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		filter := bson.M{"_id": oid}

		result := collection.FindOne(ctx, filter)

		if result.Err() != nil {
			return c.JSON(http.StatusNotFound, result.Err().Error())
		}

		err = result.Decode(&m)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, &m)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func connectToDatabase(host, port, database, collection string) *mongo.Collection {
	ctx := context.TODO()
	mongoDbUrl := fmt.Sprintf("mongodb://%s:%s", host, port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUrl))

	if err != nil {
		panic(err)
	}

	db := client.Database(database)

	c := db.Collection(collection)

	return c
}
