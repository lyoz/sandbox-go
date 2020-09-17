package main

// SES constants
var SES = struct {
	Username string
	Password string
	Host     string
	Port     int
	From     string
}{
	Username: "<access key>",
	Password: "<secret access key>",
	Host:     "email-smtp.ap-northeast-1.amazonaws.com",
	Port:     587,
	From:     "<from email address>",
}
