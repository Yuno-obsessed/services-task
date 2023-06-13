package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"services-task/receiver/model"
)

type MessageService struct {
	Conn *mongo.Collection
}

func NewMessageService() (*MessageService, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://sanity:wordpass@localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	collection := client.Database("services").Collection("messages")
	return &MessageService{
		collection,
	}, nil
}

func (r *MessageService) SaveMessage(ctx context.Context, model model.Message) (primitive.ObjectID, error) {
	model.Id = primitive.NewObjectID()
	res, err := r.Conn.InsertOne(ctx, model)
	//return fmt.Sprintf("%v", res.InsertedID), err
	return res.InsertedID.(primitive.ObjectID), err
}

func (r *MessageService) GetMessage(ctx context.Context, id primitive.ObjectID) (model.Message, error) {
	var result model.Message
	err := r.Conn.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&result)
	return result, err
}

func (r *MessageService) GetAllMessages(ctx context.Context) ([]model.Message, error) {
	var result []model.Message
	curs, err := r.Conn.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer curs.Close(ctx)
	if err := curs.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, nil
}
