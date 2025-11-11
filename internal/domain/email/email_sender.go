package email

type Sender interface {
	SendVerificationEmail(toEmail, token string) error
}
