package interfaces

import (
	"context"
	"services-task/receiver/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository struct {
	*mongo.Collection
}

func (r MessageRepository) SaveMessage(ctx context.Context, model *model.Message) (string, error) {
	res, err := r.InsertOne(ctx, model)
	return res.InsertedID.(string), err
}
