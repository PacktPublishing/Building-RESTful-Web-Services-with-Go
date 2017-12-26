package helpers

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// EncryptService is a blueprint for our service
type EncryptService interface {
	Encrypt(context.Context, string, string) (string, error)
	Decrypt(context.Context, string, string) (string, error)
}

// MakeEncryptEndpoint forms endpoint for request/response of encrypt function
func MakeEncryptEndpoint(svc EncryptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EncryptRequest)
		message, err := svc.Encrypt(ctx, req.Key, req.Text)
		if err != nil {
			return EncryptResponse{message, err.Error()}, nil
		}
		return EncryptResponse{message, ""}, nil
	}
}

// MakeDecryptEndpoint forms endpoint for request/response of decrypt function
func MakeDecryptEndpoint(svc EncryptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DecryptRequest)
		text, err := svc.Decrypt(ctx, req.Key, req.Message)
		if err != nil {
			return DecryptResponse{text, err.Error()}, nil
		}
		return DecryptResponse{text, ""}, nil
	}
}
