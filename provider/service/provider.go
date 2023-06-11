package service

import (
	"context"
	"crypto/rand"
	random "math/rand"
	"services-task/pkg/servicespb"
	"time"
)

type ProviderServer struct {
	servicespb.UnimplementedProviderServer
}

func NewProviderServer() *ProviderServer {
	return &ProviderServer{}
}

func (s *ProviderServer) Provide(ctx context.Context, req *servicespb.SymbolsRequest) (*servicespb.SymbolsResponse, error) {
	randomString, err := generateRandomString()

	response := &servicespb.SymbolsResponse{
		Symbols:       randomString,
		DateGenerated: time.Now().Unix(),
	}

	return response, err
}

var Random = random.New(random.NewSource(time.Now().UnixNano()))

func generateRandomString() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	length := Random.Intn(2000-1000+1) + 1000

	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		randomBytes[i] = charset[randomBytes[i]%byte(len(charset))]
	}

	return string(randomBytes), nil
}
