package main

import (
	"log"
	"mime"
	"os"
	"strconv"

	"github.com/wneessen/go-mail"
)

func main() {
	msg := mail.NewMsg()
	host := os.Getenv("SMTP_HOST")
	if host == "" {
		log.Fatal("SMTP_HOST required")
	}

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	if err := msg.From("hoge@example.test"); err != nil {
		log.Fatal(err)
	}
	if err := msg.To("fuga@example.test"); err != nil {
		log.Fatal(err)
	}

	msg.Subject(mime.BEncoding.Encode("UTF-8", "こんにちはこんにちは"))
	msg.SetBodyString(mail.TypeTextPlain, "ようこそこんにちは")

	c, err := mail.NewClient(host, mail.WithPort(port))
	if err != nil {
		log.Fatal(err)
	}
	if err := c.DialAndSend(msg); err != nil {
		log.Fatal(err)
	}
}
