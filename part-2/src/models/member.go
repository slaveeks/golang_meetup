package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type member struct {
	Id    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

type MemberModel struct {
	collection *mongo.Collection
	ctx        context.Context
}

func (mm *MemberModel) Create(name, email string) (*member, error) {
	m := member{
		Name:  name,
		Email: email,
	}

	_, err := mm.collection.InsertOne(mm.ctx, m)

	if err != nil {
		log.Println("Error, while inserting data")
		return nil, err
	}

	return &m, nil
}

func (mm *MemberModel) FindById(id string) (*member, error) {
	var m *member

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Error while converting hex to ObjectId")

		return nil, err
	}

	filter := bson.M{"_id": oid}

	result := mm.collection.FindOne(mm.ctx, filter)

	if result.Err() != nil {
		log.Printf("Error to find data by id: %s", result.Err())

		return nil, err
	}

	err = result.Decode(&m)

	if err != nil {
		log.Println("Error while decoding data from database")

		return nil, err
	}

	return m, nil
}

func (mm *MemberModel) FindAll() ([]*member, error) {
	var m []*member
	cursor, _ := mm.collection.Find(mm.ctx, bson.D{})

	err := cursor.All(mm.ctx, &m)

	if err != nil {
		log.Println("Error while decoding data from database")

		return nil, err
	}

	return m, nil
}

func (mm *MemberModel) Delete(id string) (*member, error) {
	var m *member

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Error while converting hex to ObjectId")

		return nil, err
	}

	filter := bson.M{"_id": oid}

	result := mm.collection.FindOneAndDelete(mm.ctx, filter)

	if result.Err() != nil {
		log.Printf("Error to find data by id: %s\n", result.Err())

		return nil, err
	}

	err = result.Decode(&m)

	if err != nil {
		log.Println("Error while decoding data from database")

		return nil, err
	}

	return m, nil
}

func CreateMemberModel(collection string, database *mongo.Database, ctx context.Context) *MemberModel {

	return &MemberModel{
		database.Collection(collection),
		ctx,
	}
}
