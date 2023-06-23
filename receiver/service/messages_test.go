package service

import (
	"context"
	"github.com/Yuno-obsessed/services-task/receiver/dto"
	"github.com/Yuno-obsessed/services-task/receiver/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var ctx = context.Background()
var service *MessageService
var model1 model.Message
var model2 model.Message

func testInit() {
	service, _ = NewMessageService()
	model1 = genModel("logs test1")
	model2 = genModel("logs test2")
}

func genModel(logs string) model.Message {
	return model.Message{
		Logs:      logs,
		CreatedAt: time.Now().UTC().Truncate(time.Second),
		StoredAt:  time.Now().UTC().Truncate(time.Second),
	}
}

func TestMessageService_SaveMessage(t *testing.T) {
	testInit()

	id, err := service.SaveMessage(ctx, model1)
	model1.Id = id
	assert.NoError(t, err)

	res, err := service.GetMessage(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, model1, res)
}

func TestMessageService_GetMessage(t *testing.T) {
	testInit()

	id, err := service.SaveMessage(ctx, model1)
	model1.Id = id
	assert.NoError(t, err)

	res, err := service.GetMessage(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, model1, res)
}

func TestMessageService_GetAllMessages(t *testing.T) {
	testInit()

	resultsAll, err := service.GetAllMessages(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resultsAll)
}

func TestMessageService_GetWithFilters(t *testing.T) {
	testInit()

	var pageSize int64 = 80
	var lengthLess int64 = 100
	//var lengthGreater int64 = 20
	filters := dto.Filters{
		Page:       0,
		PageSize:   pageSize,
		LengthLess: lengthLess,
		//LengthGreater:       lengthGreater,
		DateGeneratedAfter:  time.Date(2020, 4, 15, 0, 0, 0, 0, time.Local).Unix(),
		DateGeneratedBefore: time.Date(2023, 10, 4, 0, 0, 0, 0, time.Local).Unix(),
		Match:               "empty",
	}

	resultWithFilter, err := service.GetWithFilters(ctx, filters)
	assert.NoError(t, err)
	assert.NotEmpty(t, resultWithFilter)
	assert.Equal(t, pageSize, int64(len(resultWithFilter)))
	//for _, v := range resultWithFilter {
	//	assert.True(t, int64(len(v.Logs)) > lengthGreater)
	//}
}

func TestMessageService_Delete(t *testing.T) {
	testInit()

	id, err := service.SaveMessage(ctx, model2)
	assert.NoError(t, err)

	res, err := service.GetMessage(ctx, id)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)

	err = service.Delete(ctx, id)
	assert.NoError(t, err)

	res, err = service.GetMessage(ctx, id)
	assert.Error(t, err)
	assert.Empty(t, res)
}
