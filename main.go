package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

func sendMail() {
	m := gomail.NewMessage()

	m.SetHeader("From", SES.From)
	m.SetHeader("To", SES.To)
	m.SetHeader("Subject", "subject test (from smtp by gomail)")
	m.SetBody("text/plain", "This is the message (from smtp by gomail)")

	d := gomail.NewDialer(SES.Host, SES.Port, SES.Username, SES.Password)
	err := d.DialAndSend(m)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// sendMail()
	sendMail2()
}
