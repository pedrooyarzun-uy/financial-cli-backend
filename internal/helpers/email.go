package helpers

import (
	n "net/mail"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func ValidateEmail(email string) bool {
	_, err := n.ParseAddress(email)
	return err == nil
}

func SendMail(toEmail, toName, subject, body string) error {
	from := mail.NewEmail("Financial CLI", "me@pedrooyarzun.xyz")
	to := mail.NewEmail(toName, toEmail)

	message := mail.NewSingleEmail(from, subject, to, body, body)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	_, err := client.Send(message)

	return err
}
