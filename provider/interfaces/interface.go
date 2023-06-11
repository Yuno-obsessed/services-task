package server

import (
	"context"
	"services-task/pkg/servicespb"
)

type providerService interface {
	Provide(context.Context, *servicespb.SymbolsRequest) (*servicespb.SymbolsResponse, error)
}
