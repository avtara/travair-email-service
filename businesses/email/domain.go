package email

type Domain struct {
	EmailReceive 	string
	Token	string
}

type Service interface {
	SendNotificationCreateUsers(domain *Domain)
}

