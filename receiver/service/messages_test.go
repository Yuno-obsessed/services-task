package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"services-task/receiver/model"
	"testing"
	"time"
)

func TestMessageService(t *testing.T) {
	services, err := NewMessageService()
	assert.NoError(t, err)
	m := model.Message{
		Text:      "some text again test",
		CreatedAt: time.Now().UTC(),
		StoredAt:  time.Now().UTC(),
	}
	id, err := services.SaveMessage(context.Background(), m)
	m.Id = id
	assert.NoError(t, err)
	res, err := services.GetMessage(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, m, res)

	resultsAll, err := services.GetAllMessages(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, resultsAll)
}
