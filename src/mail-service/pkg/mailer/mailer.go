package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"igbot.com/sendmail/cmd/server/configs"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

type MailTemplate struct {
	HtmlTemplate bool
	Lang         string
	TemplateName string
	Params       map[string]interface{}
}
type Mailer struct {
	Host       string
	Port       string
	Encryption string
	Username   string
	Password   string
	Template   MailTemplate
}

func NewMailer(mailer Mailer) *Mailer {
	return &Mailer{
		Host:       mailer.Host,
		Port:       mailer.Port,
		Encryption: mailer.Encryption,
		Username:   mailer.Username,
		Password:   mailer.Password,
		Template:   mailer.Template,
	}
}

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}
func (mailer *Mailer) Sender(subject string, from string, to []string, cc []string, bcc []string) error {
	fmt.Println(mailer.Template)
	// Email content
	body, err := mailer.RenderMailTemplate()

	if err != nil {
		log.Fatalln("Can't render mail template: ", err)
		return err
	}
	// SMTP authentication credentials (not needed for MailHog)
	//auth := smtp.PlainAuth("", "", "", smtpHost)
	auth := unencryptedAuth{
		smtp.PlainAuth(
			"",
			mailer.Username,
			mailer.Password,
			mailer.Host,
		),
	}

	// Create the email message
	htmlSupport := ""
	if mailer.Template.HtmlTemplate {
		htmlSupport = "Content-Type: text/html; charset=\"UTF-8\" \r\n"
	}

	msg := []byte("To: " + strings.Join(to, ",") + "\r\n" +
		"Cc: " + strings.Join(cc, ",") + "\r\n" +
		"Bcc: " + strings.Join(bcc, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0 \r\n" +
		htmlSupport +
		"\r\n" +
		body + "\r\n")

	// Connect to the MailHog SMTP server
	err = smtp.SendMail(fmt.Sprintf("%s:%s", mailer.Host, mailer.Port), auth, from, append(append(to, cc...), bcc...), msg)
	if err != nil {
		panic(err)
	}
	return nil
}
func (mailer *Mailer) RenderMailTemplate() (string, error) {
	config := configs.LoadConfig()
	templatePath := fmt.Sprintf("%s/%s/%s.html", config.TEMPLATEFOLDERPATH, mailer.Template.Lang, mailer.Template.TemplateName)

	filePath, _ := filepath.Abs(templatePath)

	htmlTemplate, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	content, err := GenerateHtml(mailer.Template.Params, string(htmlTemplate))
	if err != nil {
		return "", err
	}
	return content, nil
}

func GenerateHtml(data map[string]interface{}, htmlTemplate string) (string, error) {
	buf := new(bytes.Buffer)

	tmpl := template.New("mail_tmp")
	tmpl, err := tmpl.Parse(htmlTemplate)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
