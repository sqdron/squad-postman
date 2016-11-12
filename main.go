package main

import (
	"github.com/sqdron/squad"
	"github.com/sqdron/squad-postman/postman"
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Options struct {
	MailHost     string
	MailPort     int
	MailUser     string
	MailPassword string
}

func main() {
	o := &Options{}
	client := squad.Client("squad.postman", o)
	d := gomail.NewDialer(o.MailHost, o.MailPort, o.MailUser, o.MailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	controller := postman.PostmanController(d)
	client.Api.Controller(controller)
	<-client.Run()
}
