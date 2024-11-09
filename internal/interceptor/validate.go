package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

// ValidateInterceptor - проверяет получаемый данные в запросах для gRPC
func ValidateInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// протовалидатор
	if val, ok := req.(validator); ok {
		if err := val.Validate(); err != nil {
			return nil, err
		}
	}
	// кастомный валидатор
	if err := ValidateRequest(req); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
