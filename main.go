package main

import (
	"fmt"
	"log"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

// username = <access key>
// password = <secret access key>
// host     = email-smtp.ap-northeast-1.amazonaws.com
// port     = 587
// from     = from email address
// to       = to email address

func sendMail() {
	auth := smtp.PlainAuth("", username, password, host)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: subject test (from smtp)\r\n" +
		"\r\n" +
		"This is the message (from smtp)\r\n")

	err := smtp.SendMail(fmt.Sprintf("%s:%d", host, port), auth, from, []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func sendMail2() {
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "subject test (from smtp by gomail)")
	m.SetBody("text/plain", "This is the message (from smtp by gomail)")

	d := gomail.NewDialer(host, port, username, password)
	err := d.DialAndSend(m)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// sendMail()
	sendMail2()
}
