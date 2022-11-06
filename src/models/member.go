package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		return nil, fmt.Errorf("failed to create member: %v", err)
	}

	return &m, nil
}

func (mm *MemberModel) FindById(id string) (*member, error) {
	var m *member

	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}

	result := mm.collection.FindOne(mm.ctx, filter)

	if result.Err() != nil {
		return nil, fmt.Errorf("error to find data by id: %s", id)
	}

	err := result.Decode(&m)

	if err != nil {
		return nil, fmt.Errorf("error while decoding data from database: %s", err)
	}

	return m, nil
}

func (mm *MemberModel) FindAll() ([]*member, error) {
	var m []*member
	cursor, _ := mm.collection.Find(mm.ctx, bson.D{})

	err := cursor.All(mm.ctx, &m)

	if err != nil {
		return nil, fmt.Errorf("error while decoding data from database: %s", err)
	}

	return m, nil
}

func (mm *MemberModel) Delete(id string) (*member, error) {
	var m *member

	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}

	result := mm.collection.FindOneAndDelete(mm.ctx, filter)

	if result.Err() != nil {
		return nil, fmt.Errorf("error to find and delete data by id: %s", id)
	}

	err := result.Decode(&m)

	if err != nil {
		return nil, fmt.Errorf("error while decoding data from database: %s", err)
	}

	return m, nil
}

func CreateMemberModel(collection string, database *mongo.Database, ctx context.Context) *MemberModel {

	return &MemberModel{
		database.Collection(collection),
		ctx,
	}
}
