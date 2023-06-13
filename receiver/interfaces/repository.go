package interfaces

import (
	"context"
	"services-task/receiver/model"
)

type MessagesRepository interface {
	SaveMessage(context.Context, *model.Message)
}