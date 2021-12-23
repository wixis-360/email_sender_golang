package config

import (
	"Email_Sender_Using_GoLang/models"
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
	"html/template"
	"log"
	"os"
)

var sender="type sender email here..." //type sender email
var receiver="type receiver email here..." // type receiver email
//password is hidden in the .env file

func setHTMLBody(user models.User) string {
	//set html template
	t := template.New("index.html")

	var err error
	//parse html file you want
	t, err = t.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, user); err != nil {
		log.Println(err)
	}
	// return tpl as a string
	htmlBody:=tpl.String()
	return htmlBody
}

func configEmail(password string, htmlBody string) (*gomail.Dialer, *gomail.Message) {

	//create New Message Function
	msg := gomail.NewMessage()

	// Set E-Mail sender
	msg.SetHeader("From", sender)

	// Set E-Mail receivers
	msg.SetHeader("To", receiver)

	// Set E-Mail subject
	msg.SetHeader("Subject", "This is a test email from example go project")

	// Set E-Mail body. You can set plain text or html with text/html
	msg.SetBody("text/html", htmlBody )

	// attach file you want
	msg.Attach("assets/wixis360.png")

	// Settings for SMTP server
	dialer := gomail.NewDialer("smtp.gmail.com", 587, sender, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	//return msg and dialer
	return dialer,msg
}

func SendMail(user models.User) {
	//get environment variable
	godotenv.Load()
	password :=os.Getenv("Password")

	//set html page
	htmlBody  := setHTMLBody(user)

	//set email
	dialer,msg:=configEmail(password,htmlBody)

	// Now send E-Mail
	if err := dialer.DialAndSend(msg); err != nil {
		fmt.Println(err)
		panic(err)
	}else {
		fmt.Println("Send Success...")
	}
}
