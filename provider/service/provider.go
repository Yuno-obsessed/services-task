package service

import (
	"context"
	"crypto/rand"
	"github.com/Yuno-obsessed/services-task/pkg/servicespb"
	random "math/rand"
	"time"
)

type ProviderServer struct {
	servicespb.UnimplementedProviderServer
}

func NewProviderServer() *ProviderServer {
	return &ProviderServer{}
}

func (s *ProviderServer) Provide(ctx context.Context, req *servicespb.ProvideLogsRequest) (*servicespb.ProvideLogsResponse, error) {
	randomString, err := generateRandomString()

	response := &servicespb.ProvideLogsResponse{
		Logs:          randomString,
		DateGenerated: time.Now().Unix(),
	}

	return response, err
}

var Random = random.New(random.NewSource(time.Now().UnixNano()))

func generateRandomString() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	length := Random.Intn(25000-10000+1) + 10000

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
