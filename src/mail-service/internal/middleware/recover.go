package middleware

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func RecoverInterceptor(p interface{}) error {
	log.Printf("Recovered panic: %v", p)
	return status.Errorf(codes.Internal, "Internal server error")
}
