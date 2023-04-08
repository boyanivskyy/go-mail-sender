package mailSender

import (
	"fmt"
	"log"
	"os"

	"github.com/go-mail/mail"
	"github.com/joho/godotenv"
)

func InitEnvVars() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Can't load .env file")
	}
}

func GetEnvVar(name string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Can't find %s env variable", name)
	}

	return val
}

type Sender struct {
	Email      string
	Password   string
	SmtpServer string
}

func InitMailSender() {
	InitEnvVars()

	sEmail := GetEnvVar("MAIL_USERNAME")
	sPass := GetEnvVar("MAIL_PASSWORD")
	smtpServer := GetEnvVar("SMTP_SERVER")

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
