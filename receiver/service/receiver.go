package service

import (
	"context"
	"log"
	"services-task/pkg/servicespb"
	"services-task/receiver/model"
	"time"
)

type ReceiverService struct {
	servicespb.UnimplementedReceiverServer
	Service *MessageService
}

func NewReceiverService() *ReceiverService {
	message, err := NewMessageService()
	if err != nil {
		log.Fatalln(err)
	}
	return &ReceiverService{
		Service: message,
	}
}

func (s *ReceiverService) Receive(ctx context.Context, request *servicespb.SymbolsResponse) (*servicespb.ProcessedSymbols, error) {

	m := model.Message{
		Text:      request.Symbols,
		CreatedAt: time.Unix(request.DateGenerated, 0).Local(),
		StoredAt:  time.Now().Local(),
	}

	_, err := s.Service.SaveMessage(ctx, m)
	if err != nil {
		return nil, err
	}
	response := &servicespb.ProcessedSymbols{
		Symbols:       request.Symbols,
		Length:        int64(len(request.Symbols)),
		DateGenerated: request.DateGenerated,
		DateSaved:     m.StoredAt.Unix(),
		//
	}

	return response, nil
}
