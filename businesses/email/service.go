package email

import (
	"fmt"
	"github.com/avtara/travair-email-service/app/config"
	"gopkg.in/gomail.v2"
	"log"
)

type emailService struct {
	configEmail *config.EmailEnv
}

func NewEmailService(conf *config.EmailEnv) Service {
	return &emailService{
		configEmail: conf,
	}
}

func (es *emailService) SendNotificationCreateUsers(domain *Domain) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", es.configEmail.ConfigSenderName)
	mailer.SetHeader("To", domain.EmailReceive)
	mailer.SetHeader("Subject", "Thank You for Registering")
	mailer.SetBody("text/html", fmt.Sprintf("Hello, <b>have a nice day %s</b>", domain.Token))

	dialer := gomail.NewDialer(
		es.configEmail.ConfigSmtpHost,
		es.configEmail.ConfigSmtpPort,
		es.configEmail.ConfigAuthEmail,
		es.configEmail.ConfigAuthPassword,
	)
	fmt.Println(es.configEmail.ConfigSenderName)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
}