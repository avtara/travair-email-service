package main

import (
	"github.com/avtara/travair-email-service/app/config"
	"github.com/avtara/travair-email-service/businesses/email"
	_emailController "github.com/avtara/travair-email-service/controllers/email"
)

func main() {
	var(
		emailEnv = config.GetEmailEnv()
		amqpConn = config.CreateConnectionAMQP()
		emailService = email.NewEmailService(emailEnv)
		emailController = _emailController.NewUserController(emailService, amqpConn)
	)

	emailController.Listener()
}

