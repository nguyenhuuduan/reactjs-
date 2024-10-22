// utils/email.go
package utils

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

// Hàm gửi OTP qua Gmail
func SendOTP(email string, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "duanlovetuong@gmail.com") // Thay bằng Gmail của bạn
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	// Cấu hình SMTP của Gmail
	d := gomail.NewDialer("smtp.gmail.com", 587, "duanlovetuong@gmail.com", "jlvuvafukbgckwsp") // Thay "your_app_password" bằng mật khẩu ứng dụng của bạn

	// Gửi email
	err := d.DialAndSend(m)
	if err != nil {
		log.Println("Failed to send OTP:", err)
		return err
	}

	log.Println("OTP sent to:", email)
	return nil
}
