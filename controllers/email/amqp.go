package email

import (
	"encoding/json"
	"fmt"
	"github.com/avtara/travair-email-service/businesses/email"
	"github.com/avtara/travair-email-service/controllers/email/request"
	"github.com/streadway/amqp"
	_ "github.com/streadway/amqp"
)

type userController struct {
	emailService email.Service
	amqpConn *amqp.Channel
}

func NewUserController(es email.Service, ac *amqp.Channel) *userController {
	return &userController{
		emailService: es,
		amqpConn: ac,
	}
}


func (uc *userController) Listener() {
	msgs, _ := uc.amqpConn.Consume(
		"travair:email",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var data request.QueueRequest
			json.Unmarshal(d.Body, &data)

			uc.emailService.SendNotificationCreateUsers(request.ToDomain(&data))
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}
