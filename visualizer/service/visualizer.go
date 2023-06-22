package service

import (
	"context"
	"services-task/pkg/servicespb"
)

type VisualizerService struct {
	servicespb.UnimplementedVisualizerServer
}

func NewVisualizerService() *VisualizerService {
	return &VisualizerService{}
}

func (v *VisualizerService) Visualize(ctx context.Context, request *servicespb.VisualizeRequest) (*servicespb.VisualizeResponse, error) {
	response := servicespb.VisualizeResponse{}

	for _, v := range request.Logs.Logs {
		responseElement := servicespb.VisualizeInfo{
			Logs:               v.Logs,
			Length:             v.Length,
			DateGenerated:      v.DateGenerated,
			DateSaved:          v.DateSaved,
			DiffGeneratedSaved: v.DateGenerated - v.DateSaved,
		}
		response.Info = append(response.Info, &responseElement)
	}
	return &response, nil
}
