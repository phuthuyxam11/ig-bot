package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"igbot.com/sendmail/api/pb"
)

func (server *MailSenderServer) SendMailRecoveryPassword(context.Context, *pb.MailRecoveryAccRequest) (*pb.MailRecoveryAccResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMailRecoveryPassword not implemented")
}
