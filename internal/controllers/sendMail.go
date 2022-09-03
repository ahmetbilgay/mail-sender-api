package controllers

import (
	"blgy-mail-sender/internal/models"
	"crypto/tls"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
	"log"
)

func SendMail(c *fiber.Ctx) error {
	// IMPORT MAILMODEL
	mail := new(models.MailModel)

	// GET POST DATA
	if err := c.BodyParser(&mail); err != nil {
		return err
	}

	// INITIALIZE GOMAIL
	msg := gomail.NewMessage()

	// GOMAIL SET HEADER
	msg.SetHeader("From", mail.MsgSender)
	msg.SetHeader("To", mail.MsgTo)
	msg.SetHeader("Subject", mail.Subject)

	// GOMAIL SET BODY
	msg.SetBody("text/plain", mail.BodyMessage)

	// GOMAIL DIALER
	x := gomail.NewDialer("smtp.gmail.com", 587, mail.MsgSender, mail.Password)

	// GOMAIL TLS CONFIG
	x.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// GOMAIL ERROR HANDLING
	if err := x.DialAndSend(msg); err != nil {
		log.Println(err)
		panic(err)
	}

	// RETURN RESPONSE DATA
	return c.SendStatus(200)
}
