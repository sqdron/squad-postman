package postman

import "gopkg.in/gomail.v2"

type postman struct {
	mail *gomail.Dialer
}

func PostmanController(mail *gomail.Dialer) *postman {
	return &postman{mail}
}
