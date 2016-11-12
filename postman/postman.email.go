package postman

import "gopkg.in/gomail.v2"

type MailRequest struct {
	From    string
	To      string
	Subject string
	Body    string
	CC      string
}

func (c *postman) SendMail(r *MailRequest) error {
	m := gomail.NewMessage()
	m.SetHeader("From", r.From)
	m.SetHeader("To", r.To)
	m.SetHeader("Subject", r.Subject)
	m.SetBody("text/html", r.Body)

	return c.mail.DialAndSend(m);
}
