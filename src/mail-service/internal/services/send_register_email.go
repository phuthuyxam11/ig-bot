package services

import (
	"context"
	"igbot.com/sendmail/api/pb"
	"igbot.com/sendmail/pkg/helpers"
	"igbot.com/sendmail/pkg/mailer"
	"igbot.com/sendmail/pkg/mailer/mail_server"
)

func (server *MailSenderServer) SendMailVerifyAcc(ctx context.Context, req *pb.MailVerifyAccRequest) (*pb.MailVerifyAccResponse, error) {

	template := mailer.MailTemplate{
		Lang:         req.EmailInfo.Lang,
		TemplateName: "acc_register",
		Params:       map[string]interface{}{"register_url": req.ActiveUrl},
		HtmlTemplate: true,
	}
	mailerTemplate := mail_server.MailHogServer(template)
	mailFactory := mailer.NewMailFactory(mailerTemplate)
	err := mailFactory.SendMail(req.EmailInfo.Subject,
		req.EmailInfo.From,
		helpers.MapToStringSlice(req.EmailInfo.To),
		helpers.MapToStringSlice(req.EmailInfo.Cc),
		helpers.MapToStringSlice(req.EmailInfo.Bcc))
	if err != nil {
		return nil, err
	}
	res := &pb.MailVerifyAccResponse{
		Status:  "asd",
		Message: "adad",
	}
	return res, nil
}
