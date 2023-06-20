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

func (v *VisualizerService) Visualize(ctx context.Context, request *servicespb.ProcessedSymbols) (*servicespb.PresentedSymbols, error) {
	response := servicespb.PresentedSymbols{
		Symbols:         request.Symbols,
		Length:          request.Length,
		DataGenerated:   request.DateGenerated,
		DateSaved:       request.DateSaved,
		ProviderService: request.ProviderService,
		ReceiverService: request.ReceiverService,
	}
	return &response, nil
}
