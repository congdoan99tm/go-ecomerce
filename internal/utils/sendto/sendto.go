package sendto

import (
	"fmt"
	"github.com/dinos/go-ecommerce-be-api/global"
	"go.uber.org/zap"
	"net/smtp"
	"strings"
)

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPUser     = "colucnatrat92@gmail.com"
	SMTPPassword = "a"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MINE-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s", mail.Body)
	return msg
}

func SendTextEmailOTP(to []string, from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP verification",
		Body:    fmt.Sprintf("Your OTP is %s, Please enter it to verify your account.", otp),
	}
	messageMail := BuildMessage(contentEmail)
	auth := smtp.PlainAuth("", SMTPUser, SMTPPassword, SMTPHost)
	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send failed", zap.Error(err))
		return err
	}
	return nil
}
