package service

import (
	"context"
	"log"
	"services-task/pkg/servicespb"
	"services-task/receiver/dto"
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

func (s *ReceiverService) Receive(ctx context.Context, request *servicespb.ReceiveLogsRequest) (*servicespb.ResponseStatus, error) {
	var response servicespb.ResponseStatus

	m := model.Message{
		Logs:      request.Logs.Logs,
		CreatedAt: time.Unix(request.Logs.DateGenerated, 0).UTC().Truncate(time.Second),
		StoredAt:  time.Now().UTC().Truncate(time.Second),
	}

	_, err := s.Service.SaveMessage(ctx, m)
	if err != nil {
		response.Status = 404
		return &response, err
	}
	response.Status = 200

	return &response, nil
}

func (s *ReceiverService) Fetch(ctx context.Context, request *servicespb.Filters) (*servicespb.FetchResponse, error) {
	var response servicespb.FetchResponse

	var filters = dto.Filters{
		Page:                request.Page,
		PageSize:            request.PageSize,
		Match:               request.Match,
		DateGeneratedAfter:  request.DateGeneratedAfter,
		DateGeneratedBefore: request.DateGeneratedBefore,
		LengthLess:          request.LengthLess,
		LengthGreater:       request.LengthGreater,
	}

	m, err := s.Service.GetWithFilters(ctx, filters)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	logs := []*servicespb.FetchedLogs{}
	for _, v := range m {
		resLogs := servicespb.FetchedLogs{
			Logs:          v.Logs,
			Length:        int64(len(v.Logs)),
			DataGenerated: v.CreatedAt.Unix(),
			DateSaved:     v.StoredAt.Unix(),
		}
		logs = append(logs, &resLogs)
	}

	response.Logs = logs

	return &response, nil
}
