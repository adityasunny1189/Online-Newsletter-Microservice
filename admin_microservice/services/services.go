package services

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	msg = `%v
By: %v
`
)

func Mail(author, heading, content string, destEmails []string) {
	from := /*os.Getenv("MAIL")*/ "aditya9102833743@gmail.com"
	password := /*os.Getenv("PASSWD")*/ "zblbrafgrmhamime"

	toList := destEmails
	host := "smtp.gmail.com"
	port := "587"

	body := []byte("Subject: " + heading + "\r\n\r\n" + "" + fmt.Sprintf(msg, content, author))
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		log.Fatalf("error sending mail %v", err)
	}
	log.Printf("Successfully sent mail to the user")
}
