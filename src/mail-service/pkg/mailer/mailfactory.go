package mailer

type Factory interface {
	RenderMailTemplate() (string, error)
	Sender(subject string, from string, to []string, cc []string, bcc []string) error
}

type MailFactory struct {
	factory Factory
}

func NewMailFactory(factory Factory) *MailFactory {
	return &MailFactory{factory: factory}
}

func (factory *MailFactory) RenderMailTemplateImpl() (string, error) {
	return factory.factory.RenderMailTemplate()
}
func (factory *MailFactory) SendMail(subject string, from string, to []string, cc []string, bcc []string) error {
	return factory.factory.Sender(subject, from, to, cc, bcc)
}
