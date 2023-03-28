package services

import (
	"igbot.com/sendmail/api/pb"
	"igbot.com/sendmail/component"
)

type MailSenderServer struct {
	appContext component.AppContext
	pb.UnimplementedMailSenderServiceServer
}

func NewMailSenderServer(appCtx component.AppContext) *MailSenderServer {
	return &MailSenderServer{
		appContext:                           appCtx,
		UnimplementedMailSenderServiceServer: pb.UnimplementedMailSenderServiceServer{},
	}
}
