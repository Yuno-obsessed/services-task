package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"services-task/receiver/dto"
	"services-task/receiver/model"
	"time"
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

func (r *MessageService) GetWithFilters(ctx context.Context, filters dto.Filters) ([]model.Message, error) {
	var result []model.Message
	var filter bson.M
	filterLengthGreater := bson.M{}
	filterLengthLess := bson.M{}
	filterDateAfter := bson.M{}
	filterDateBefore := bson.M{}
	page := filters.Page
	pageSize := filters.PageSize
	var skip int64
	if filters.LengthLess != 0 && filters.LengthGreater != 0 {
		filters.LengthLess = 0
	}
	if filters.DateGeneratedAfter != 0 {
		filterDateAfter = bson.M{
			"created_at": bson.M{"$gt": time.Unix(filters.DateGeneratedAfter, 0)},
		}
	} else if filters.DateGeneratedBefore != 0 {
		filterDateBefore = bson.M{
			"created_at": bson.M{"$lt": time.Unix(filters.DateGeneratedBefore, 0)},
		}
	}

	if filters.LengthGreater != 0 {
		filterLengthGreater = bson.M{
			"$expr": bson.M{
				"$gt": []interface{}{
					bson.M{"$strLenCP": "$text"},
					filters.LengthGreater,
				},
			},
		}
	} else if filters.LengthLess != 0 {
		filterLengthLess = bson.M{
			"$expr": bson.M{
				"$lt": []interface{}{
					bson.M{"$strLenCP": "$text"},
					filters.LengthLess,
				},
			},
		}
	}

	findOptions := options.Find()
	if filters.PageSize == 0 {
		pageSize = 20
		skip = (page - 1) * pageSize
	}

	findOptions.SetSkip(skip)
	findOptions.SetLimit(pageSize)

	if filters.DateGeneratedAfter != 0 {
		filter = bson.M{
			"$and": []bson.M{
				{
					"text": bson.M{
						"$ne":      "",
						"$regex":   filters.Match,
						"$options": "i",
					},
				},
				filterDateAfter,
				filterDateBefore,
				filterLengthGreater,
				filterLengthLess,
			},
		}
	}

	res, err := r.Conn.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	for res.Next(context.Background()) {
		var message model.Message
		if err := res.Decode(&message); err != nil {
			fmt.Println("Failed to decode document:", err)
			return nil, err
		}
		fmt.Println(res)
		result = append(result, message)
	}

	if err := res.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return nil, err
	}
	return result, nil
}
