package mail_server

import "igbot.com/sendmail/pkg/mailer"

func MailHogServer(template mailer.MailTemplate) *mailer.Mailer {
	return &mailer.Mailer{
		Host:       "mailhog",
		Port:       "1025",
		Encryption: "",
		Username:   "",
		Password:   "",
		Template:   template,
	}
}
