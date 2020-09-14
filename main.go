package main

import (
	"fmt"
	"log"
	"net/smtp"
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

func main() {
	sendMail()
}
