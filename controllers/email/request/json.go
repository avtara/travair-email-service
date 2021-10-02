package request

import "github.com/avtara/travair-email-service/businesses/email"

type QueueRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Type  string `json:"type"`
}

func ToDomain(data *QueueRequest) *email.Domain {
	return &email.Domain{
		EmailReceive: data.Email,
		Token: data.Token,
	}
}