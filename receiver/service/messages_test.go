package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"services-task/receiver/dto"
	"services-task/receiver/model"
	"testing"
	"time"
)

func TestMessageService(t *testing.T) {
	services, err := NewMessageService()
	assert.NoError(t, err)
	m := model.Message{
		Logs:      "some text again test",
		CreatedAt: time.Now().UTC().Truncate(time.Second),
		StoredAt:  time.Now().UTC().Truncate(time.Second),
	}
	id, err := services.SaveMessage(context.Background(), m)
	m.Id = id
	assert.NoError(t, err)
	res, err := services.GetMessage(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, m, res)

	resultsAll, err := services.GetAllMessages(context.Background())
	assert.NoError(t, err)
	//fmt.Println(resultsAll)
	assert.NotEmpty(t, resultsAll)

	var pageSize int64 = 10
	var lengthLess int64 = 100
	var lengthGreater int64 = 20
	filters := dto.Filters{
		Page:                0,
		PageSize:            pageSize,
		LengthLess:          lengthLess,
		LengthGreater:       lengthGreater,
		DateGeneratedAfter:  time.Date(2020, 4, 15, 0, 0, 0, 0, time.Local).Unix(),
		DateGeneratedBefore: time.Date(2023, 10, 4, 0, 0, 0, 0, time.Local).Unix(),
		//Match:               "some",
	}
	resultWithFilter, err := services.GetWithFilters(context.Background(), filters)
	assert.NoError(t, err)
	//fmt.Println(resultWithFilter)
	assert.NotEmpty(t, resultWithFilter)
	//assert.Equal(t, pageSize, int64(len(resultWithFilter)))
	for _, v := range resultWithFilter {
		assert.True(t, int64(len(v.Logs)) > lengthGreater)
	}
}
