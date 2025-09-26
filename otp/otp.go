package otp

import (
	"errors"
	"fmt"
	"math/rand"

	"time"

	"gopkg.in/gomail.v2"
)

// var otpStore = make(map[string]string)

func SendOTP(email string) (string, error) {

	// Generate OTP
	otp := GenerateOTP()
	err := SendOTPEmail(email, otp)
	if err != nil {
		return "", errors.New("unable to send email")

	}
	return otp, nil

}
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(999999))
}

func SendOTPEmail(email, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "abce10981@gmail.com")
	m.SetHeader("To", "ankita.s2805@gmail.com")
	m.SetHeader("Subject", "Your OTP")
	m.SetBody("text/html", fmt.Sprintf("Your OTP for registration is:  %s!", otp))

	d := gomail.NewDialer("smtp.example.com", 587, "ankita.s2805@gmail.com", "oqfp xjzf prxp cjzc")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil
}


