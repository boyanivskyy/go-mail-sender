package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-mail/mail"
	"github.com/joho/godotenv"
)

func initEnvVars() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Can't load .env file")
	}
}

func getEnvVar(name string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Can't find sender email")
	}

	return val
}

type Sender struct {
	Email      string
	Password   string
	SmtpServer string
}

func main() {
	initEnvVars()

	sEmail := getEnvVar("MAIL_USERNAME")
	sPass := getEnvVar("MAIL_PASSWORD")
	smtpServer := getEnvVar("SMTP_SERVER")

	sender := Sender{
		Email:      sEmail,
		Password:   sPass,
		SmtpServer: smtpServer,
	}

	receiver := ""

	fmt.Print("Receiver -> ")
	_, _ = fmt.Scanln(&receiver)
	if receiver == "" {
		log.Fatalf("Empty receiver email")
	}

	mailContent := ""

	fmt.Print("Content -> ")
	_, _ = fmt.Scanln(&mailContent)
	if mailContent == "" {
		log.Fatalf("Empty mail content")
	}

	m := mail.NewMessage()
	m.SetHeader("From", sender.Email)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", mailContent)

	d := mail.NewDialer(sender.SmtpServer, 587, sender.Email, sender.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	log.Printf("Mail has been sent to %s, with content\n %s", receiver, mailContent)
}
