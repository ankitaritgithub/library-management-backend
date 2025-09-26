package otp

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendOtp2(email, otp string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "ankita.s2805@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body"+otp)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "ankita.s2805@gmail.com", "oqfp xjzf prxp cjzc")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

}
